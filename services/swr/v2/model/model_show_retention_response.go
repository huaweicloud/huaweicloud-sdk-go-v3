package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRetentionResponse struct {

	// 回收规则匹配策略，or
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm"`

	// ID
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 镜像老化规则
	Rules *[]Rule `json:"rules,omitempty" xml:"rules"`

	// 保留字段
	Scope          *string `json:"scope,omitempty" xml:"scope"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRetentionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRetentionResponse struct{}"
	}

	return strings.Join([]string{"ShowRetentionResponse", string(data)}, " ")
}
