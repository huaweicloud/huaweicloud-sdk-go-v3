package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 可选。SQL作业。
type SqlJob struct {
	// 作业的SQL。

	Sql string `json:"sql"`
}

func (o SqlJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJob struct{}"
	}

	return strings.Join([]string{"SqlJob", string(data)}, " ")
}
