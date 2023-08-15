package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeApiVersionV2Request Request Object
type ChangeApiVersionV2Request struct {

	// 实例编号
	InstanceId string `json:"instance_id"`

	// API的编号
	ApiId string `json:"api_id"`

	Body *ApiVersion `json:"body,omitempty"`
}

func (o ChangeApiVersionV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeApiVersionV2Request struct{}"
	}

	return strings.Join([]string{"ChangeApiVersionV2Request", string(data)}, " ")
}
