package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewRequestBody struct {

	// 应用列表
	ApplicationList []BatchCreateApplicationViewRequestBodyApplicationList `json:"application_list"`

	// 组件列表
	ComponentList *[]BatchCreateApplicationViewRequestBodyComponentList `json:"component_list,omitempty"`

	// 分组列表
	GroupList *[]BatchCreateApplicationViewRequestBodyGroupList `json:"group_list,omitempty"`
}

func (o BatchCreateApplicationViewRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBody", string(data)}, " ")
}
