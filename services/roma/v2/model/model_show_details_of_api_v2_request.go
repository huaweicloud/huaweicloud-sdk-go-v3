package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfApiV2Request Request Object
type ShowDetailsOfApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// API的编号
	ApiId string `json:"api_id"`
}

func (o ShowDetailsOfApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfApiV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfApiV2Request", string(data)}, " ")
}
