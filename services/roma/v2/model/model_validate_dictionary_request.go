package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateDictionaryRequest Request Object
type ValidateDictionaryRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 待校验是否重复的字典名称，精确匹配
	Name *string `json:"name,omitempty"`

	// 待校验是否重复的字典编码，精确匹配
	Code *string `json:"code,omitempty"`
}

func (o ValidateDictionaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateDictionaryRequest struct{}"
	}

	return strings.Join([]string{"ValidateDictionaryRequest", string(data)}, " ")
}
