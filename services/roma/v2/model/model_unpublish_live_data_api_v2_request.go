package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnpublishLiveDataApiV2Request Request Object
type UnpublishLiveDataApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 后端API的编号
	LdApiId string `json:"ld_api_id"`
}

func (o UnpublishLiveDataApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnpublishLiveDataApiV2Request struct{}"
	}

	return strings.Join([]string{"UnpublishLiveDataApiV2Request", string(data)}, " ")
}
