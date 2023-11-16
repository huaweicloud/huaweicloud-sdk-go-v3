package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateBussinessSceneSpecMatches struct {

	// 条件名称
	Name *string `json:"name,omitempty"`

	// 匹配的PATH
	ApiPath *interface{} `json:"apiPath,omitempty"`

	// 匹配的Headers
	Headers *interface{} `json:"headers,omitempty"`

	// 匹配的Method列表
	Method *[]string `json:"method,omitempty"`

	// 匹配的微服务名称
	ServiceName *string `json:"serviceName,omitempty"`
}

func (o CreateBussinessSceneSpecMatches) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBussinessSceneSpecMatches struct{}"
	}

	return strings.Join([]string{"CreateBussinessSceneSpecMatches", string(data)}, " ")
}
