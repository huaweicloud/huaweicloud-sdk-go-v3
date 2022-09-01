package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowLiveDataApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 后端API的编号
	LdApiId string `json:"ld_api_id" xml:"ld_api_id"`
}

func (o ShowLiveDataApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLiveDataApiV2Request struct{}"
	}

	return strings.Join([]string{"ShowLiveDataApiV2Request", string(data)}, " ")
}
