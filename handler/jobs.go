package handler

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/onyas/geekNews/model"
	"github.com/onyas/geekNews/service"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var group sync.WaitGroup

var dataFrom = []string{"V2EX", "JueJin", "EleDuck"}

//var dataFrom = []string{"JueJin"}

var jobService service.JobService

type Spider struct {
	DataType string
}

func ListJobInfos(context *gin.Context) {
	jobId, _ := strconv.Atoi(context.DefaultQuery("id", strconv.Itoa(int(^uint32(0)))))
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
	req, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36`)
	req.Header.Add("accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3`)
	req.Header.Set("X-Real-IP", randomIpAddress())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		log.Println(err)
		return []map[string]interface{}{}
	}
	defer res.Body.Close()

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
		jobId := jobIdRex.FindStringSubmatch(url)[1]
		title := selection.Find(".item_title").Text()
		author := selection.Find(".topic_info a").First().Text()
		fmt.Println(authorAvatar + " " + url + " " + title + " " + author + " " + jobId)
		allData = append(allData, map[string]interface{}{"authorAvatar": authorAvatar, "title": title, "url": "https://v2ex.com" + url, "author": author, "jobId": jobId, "dataFrom": "V2EX"})
	})
	return allData
}

func (spider Spider) GetJueJin() []map[string]interface{} {
	req, err := http.NewRequest("GET", "https://short-msg-ms.juejin.im/v1/pinList/topic?uid=&device_id=&token=&src=web&topicId=5abb61e1092dcb4620ca3322&page=0&pageSize=20&sortType=rank", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Origin", "https://juejin.im")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Referer", "https://juejin.im/pins/topic/5abb61e1092dcb4620ca3322")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,it;q=0.8,en;q=0.7")
	req.Header.Set("X-Real-IP", randomIpAddress())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jobInfos model.JueJinResponse
	if err := json.Unmarshal(body, &jobInfos); err != nil {
		log.Fatal(err)
	}

	var allData []map[string]interface{}
	for _, jobInfo := range jobInfos.D.List {
		allData = append(allData, map[string]interface{}{"authorAvatar": jobInfo.User.AvatarLarge, "title": strings.ReplaceAll(jobInfo.Content, "0x00", ""),
			"url": jobInfo.ObjectId, "author": jobInfo.User.Username, "createdAt": jobInfo.CreateAt,
			"attachMent": strings.Join(jobInfo.Pictures, ","),
			"jobId":      jobInfo.ObjectId, "dataFrom": "JueJin"})
	}
	return allData
}

func (spider Spider) GetEleDuck() []map[string]interface{} {
	req, err := http.NewRequest("GET", "https://svc.eleduck.com/api/v1/posts?category=5&tags[]=44&tags[]=10&page=1", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Origin", "https://eleduck.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	req.Header.Set("Dnt", "1")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Referer", "https://eleduck.com/?category=5&tags=44_10")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,it;q=0.8,en;q=0.7")
	req.Header.Set("If-None-Match", "W/\"29e00361915d0f9368e9eb1c01e80f3c\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jobInfos model.EleDuckResponse
	if err := json.Unmarshal(body, &jobInfos); err != nil {
		log.Fatal(err)
	}

	var allData []map[string]interface{}
	for _, jobInfo := range jobInfos.Posts {
		allData = append(allData, map[string]interface{}{"authorAvatar": jobInfo.User.AvatarUrl, "title": strings.ReplaceAll(jobInfo.Title, "0x00", ""),
			"url": "https://eleduck.com/posts/" + jobInfo.Id, "author": jobInfo.User.NickName, "createdAt": jobInfo.PublishedAt,
			"jobId": jobInfo.Id, "dataFrom": "EleDuck"})
	}
	return allData
}

func randomIpAddress() string {
	r := rand.Intn(254)
	return fmt.Sprintf("207.163.226.%d", r)
}
