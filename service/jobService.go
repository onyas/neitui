package service

import (
	"github.com/mitchellh/mapstructure"
	"github.com/onyas/geekNews/model"
	"log"
)

type JobService struct {
}

func (service *JobService) SearchJobInfos() []model.JobInfo {
	jobs := make([]model.JobInfo, 0)
	err := DbEngine.Find(&jobs)
	if nil != err {
		log.Println(err)
	}
	return jobs
}

func (service *JobService) SaveData(jobInfos []map[string]interface{}) []model.JobInfo {
	for _, jobInfo := range jobInfos {
		var job model.JobInfo
		mapstructure.Decode(jobInfo, &job)
		_, err := DbEngine.Insert(&job)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
