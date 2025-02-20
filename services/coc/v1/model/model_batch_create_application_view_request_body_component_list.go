package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchCreateApplicationViewRequestBodyComponentList struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 父节点code
	ParentName *string `json:"parent_name,omitempty"`
}

func (o BatchCreateApplicationViewRequestBodyComponentList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateApplicationViewRequestBodyComponentList struct{}"
	}

	return strings.Join([]string{"BatchCreateApplicationViewRequestBodyComponentList", string(data)}, " ")
}
