package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateLiveDataApiScriptV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 后端API的编号
	LdApiId string `json:"ld_api_id" xml:"ld_api_id"`

	Body *LdApiScriptCreate `json:"body,omitempty" xml:"body"`
}

func (o CreateLiveDataApiScriptV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLiveDataApiScriptV2Request struct{}"
	}

	return strings.Join([]string{"CreateLiveDataApiScriptV2Request", string(data)}, " ")
}
