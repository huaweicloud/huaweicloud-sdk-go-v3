package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeApiVersionV2Request struct {

	// 实例编号
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// API的编号
	ApiId string `json:"api_id" xml:"api_id"`

	Body *ApiVersion `json:"body,omitempty" xml:"body"`
}

func (o ChangeApiVersionV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeApiVersionV2Request struct{}"
	}

	return strings.Join([]string{"ChangeApiVersionV2Request", string(data)}, " ")
}
