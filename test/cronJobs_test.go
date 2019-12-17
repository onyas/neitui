package test

import (
	"fmt"
	"github.com/onyas/geekNews/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"testing"
)

func TestCronJobs(t *testing.T) {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/cronJobs", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestListCronJobs(t *testing.T) {
	router := router.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/listJobInfos?limit=3&id=10", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestRegexp(t *testing.T) {
	jobIdRex, _ := regexp.Compile(`^/t/([0-9]*)#reply[0-9]*`)
	fmt.Println(jobIdRex.FindString("/t/625140#reply66"))
	jobIds := jobIdRex.FindStringSubmatch("/t/625140#reply66")[1]
	fmt.Println(jobIds)
}

func TestInteger(t *testing.T) {
	fmt.Println(int(^uint32(0)))
	fmt.Println(strconv.Itoa(int(^uint32(0))))
}
