package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CommitsCommits struct {
	// 提交记录sha值。

	Sha string `json:"sha"`
	// 提交记录描述。

	Message string `json:"message"`
	// 合入时间。

	AuthoredDate string `json:"authored_date"`
}

func (o CommitsCommits) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitsCommits struct{}"
	}

	return strings.Join([]string{"CommitsCommits", string(data)}, " ")
}
