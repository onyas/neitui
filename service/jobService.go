package service

import (
	"github.com/mitchellh/mapstructure"
	"github.com/onyas/neitui/model"
	"log"
)

type JobService struct {
}

func (service *JobService) SearchJobInfos(offset, perPage int) []model.JobInfo {
	jobs := make([]model.JobInfo, 0)
	err := DbEngine.Desc("id").Limit(perPage, offset).Find(&jobs)
	if nil != err {
		log.Println(err)
	}
	return jobs
}

func (service *JobService) CountJobInfos() int64 {
	job := new(model.JobInfo)
	total, err := DbEngine.Where("id >?", 1).Count(job)
	if err != nil {
		log.Println(err)
		return 0
	}
	return total
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
