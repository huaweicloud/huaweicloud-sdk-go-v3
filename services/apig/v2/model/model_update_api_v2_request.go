package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateApiV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// API的编号，可通过查询API信息获取该编号。

	ApiId string `json:"api_id"`

	Body *ApiCreate `json:"body,omitempty"`
}

func (o UpdateApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiV2Request struct{}"
	}

	return strings.Join([]string{"UpdateApiV2Request", string(data)}, " ")
}
