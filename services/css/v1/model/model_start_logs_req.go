package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartLogsReq struct {
	// IAM委托。

	Agency string `json:"agency"`
	// 备份路径。

	LogBasePath string `json:"logBasePath"`
	// OBS桶。

	LogBucket string `json:"logBucket"`
}

func (o StartLogsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartLogsReq struct{}"
	}

	return strings.Join([]string{"StartLogsReq", string(data)}, " ")
}
