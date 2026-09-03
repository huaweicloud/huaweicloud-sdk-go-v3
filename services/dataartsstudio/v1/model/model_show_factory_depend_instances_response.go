package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryDependInstancesResponse Response Object
type ShowFactoryDependInstancesResponse struct {

	// 实例详情。
	DependInstancesInfo *[]ShowFactoryDependInstancesRespDependInstancesInfo `json:"depend_instances_info,omitempty"`

	// 返回的实例总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowFactoryDependInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryDependInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowFactoryDependInstancesResponse", string(data)}, " ")
}
