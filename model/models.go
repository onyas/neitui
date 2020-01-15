package model

import "time"

type Spider struct {
	DataType string
}

type JobInfo struct {
	Id           int64     `xorm:"pk autoincr bigint(20)" form:"id" json:"id,omitempty"`
	JobId        string    `xorm:"varchar(50) unique(uindex_jobid)" form:"jobid" json:"jobid"`
	Title        string    `xorm:"varchar(1000)" form:"title" json:"title"`
	Url          string    `xorm:"varchar(100)" form:"url" json:"url,omitempty"`
	Author       string    `xorm:"varchar(50)" form:"author" json:"author,omitempty"`
	AuthorAvatar string    `xorm:"varchar(100)" form:"authoravatar" json:"authorAvatar,omitempty"`
	DataFrom     string    `xorm:"varchar(50) unique(uindex_jobid)" form:"datafrom" json:"dataFrom,omitempty"`
	AttachMent   string    `xorm:"varchar(1000)" form:"attachMent" json:"attachMent"`
	CreatedAt    time.Time `xorm:"created" form:"createdAt" json:"createdAt"`
}

type JueJinUser struct {
	AvatarLarge    string `json:"avatarLarge"`
	Company        string `json:"company"`
	FollowersCount int    `json:"followersCount"`
	JobTitle       string `json:"jobTitle"`
	Username       string `json:"username"`
}

type JueJinJobInfo struct {
	Content  string     `json:"content"`
	CreateAt time.Time  `json:"createdAt"`
	ObjectId string     `json:"objectId"`
	Pictures []string   `json:"pictures,omitempty"`
	Uid      string     `json:"uid"`
	User     JueJinUser `json:"user"`
}
type JueJinData struct {
	Total int             `json:"total"`
	List  []JueJinJobInfo `json:"list"`
}

type JueJinResponse struct {
	S int        `json:"s"`
	M string     `json:"m"`
	D JueJinData `json:"d"`
}

type Result struct {
	Code    int         `form:"code" json:"code"`
	Message string      ` form:"message" json:"message"`
	Data    interface{} `json:"data" `
}

//===================================================================

type EleDuckResponse struct {
	Posts []EleDuckPost `json:"posts"`
}

type EleDuckPost struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	ModifiedAt  time.Time `json:"modified_at"`
	PublishedAt time.Time `json:"published_at"`
	User        UserInfo  `json:"user"`
}

type UserInfo struct {
	Id        string `json:"id"`
	AvatarUrl string `json:"avatar_url"`
	NickName  string `json:"nickname"`
}
