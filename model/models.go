package model

import "time"

type Spider struct {
	DataType string
}

type JobInfo struct {
	Id           int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id,omitempty"`
	JobId        int64     `xorm:"bigint(20) unique(uindex_jobid)" form:"jobId" json:"jobId"`
	Title        string    `xorm:"varchar(100)" form:"title" json:"title"`
	Url          string    `xorm:"varchar(100)" form:"url" json:"url,omitempty"`
	Author       string    `xorm:"varchar(50)" form:"author" json:"author,omitempty"`
	AuthorAvatar string    `xorm:"varchar(100)" form:"authoravatar" json:"authorAvatar,omitempty"`
	DataFrom     string    `xorm:"varchar(50) unique(uindex_jobid)" form:"datafrom" json:"dataFrom,omitempty"`
	UpdatedAt    time.Time `xorm:"updated" form:"updatedat" json:"updatedat"`
}

type Result struct {
	Code    int         `form:"code" json:"code"`
	Message string      ` form:"message" json:"message"`
	Data    interface{} `json:"data" `
}
