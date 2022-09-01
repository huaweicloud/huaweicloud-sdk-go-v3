package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ValidateRomaAppRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 应用名称，不支持模糊匹配
	Name *string `json:"name,omitempty" xml:"name"`

	// 应用ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用key
	Key *string `json:"key,omitempty" xml:"key"`
}

func (o ValidateRomaAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateRomaAppRequest struct{}"
	}

	return strings.Join([]string{"ValidateRomaAppRequest", string(data)}, " ")
}
