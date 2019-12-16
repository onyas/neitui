package handler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/onyas/geekNews/service"
	"io"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var group sync.WaitGroup
var dataFrom = []string{"V2EX"}

var jobService service.JobService

type Spider struct {
	DataType string
}

func CronJobs(context *gin.Context) {
	fmt.Println("开始抓取" + strconv.Itoa(len(dataFrom)) + "种数据类型")
	group.Add(len(dataFrom))
	var spider Spider
	for _, value := range dataFrom {
		fmt.Println("开始抓取" + value)
		spider = Spider{DataType: value}
		go ExecGetData(spider)
	}
	group.Wait()
	fmt.Print("完成抓取")

	context.JSON(http.StatusOK, gin.H{
		"message": "fetch success",
	})
}

func ExecGetData(spider Spider) {
	reflectValue := reflect.ValueOf(spider)
	dataType := reflectValue.MethodByName("Get" + spider.DataType)
	data := dataType.Call(nil)
	originData := data[0].Interface().([]map[string]interface{})
	start := time.Now()
	jobService.SaveData(originData)
	group.Done()
	seconds := time.Since(start).Seconds()
	fmt.Printf("耗费 %.2fs 秒完成抓取%s", seconds, spider.DataType)
	fmt.Println()
}

func (spider Spider) GetV2EX() []map[string]interface{} {
	url := "https://v2ex.com/go/jobs"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36`)
	request.Header.Add("accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		log.Println(err)
		return []map[string]interface{}{}
	}
	defer res.Body.Close()
	////just for test
	//html, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(html))
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	var allData []map[string]interface{}
	jobIdRex, err := regexp.Compile(`^/t/([0-9]*)#reply[0-9]*`)
	document.Find("#TopicsNode table").Each(func(i int, selection *goquery.Selection) {
		authorAvatar, _ := selection.Find("td img").Attr("src")
		url, _ := selection.Find(".topic-link").Attr("href")
		jobId, _ := strconv.Atoi(jobIdRex.FindStringSubmatch(url)[1])
		title := selection.Find(".item_title").Text()
		author := selection.Find(".topic_info a").First().Text()
		fmt.Println(authorAvatar + " " + url + " " + title + " " + author + " " + string(jobId))
		allData = append(allData, map[string]interface{}{"authorAvatar": authorAvatar, "title": title, "url": url, "author": author, "jobId": jobId, "dataFrom": "V2EX"})
	})
	return allData
}

func ListJobInfos(context *gin.Context) {
	jobId, _ := strconv.Atoi(context.DefaultQuery("jobId", "0"))
	limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
	if err != nil {
		log.Println(err)
	}
	if limit == 0 {
		limit = 10
	}
	jobInfos := jobService.SearchJobInfos(jobId, limit)
	context.JSON(http.StatusOK, gin.H{
		"jobInfos": jobInfos,
	})
}
