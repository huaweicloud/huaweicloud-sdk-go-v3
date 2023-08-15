package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateRomaAppRequest Request Object
type ValidateRomaAppRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 应用名称，不支持模糊匹配
	Name *string `json:"name,omitempty"`

	// 应用ID
	Id *string `json:"id,omitempty"`

	// 应用key
	Key *string `json:"key,omitempty"`
}

func (o ValidateRomaAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateRomaAppRequest struct{}"
	}

	return strings.Join([]string{"ValidateRomaAppRequest", string(data)}, " ")
}
