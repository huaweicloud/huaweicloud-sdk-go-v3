package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAppInstancesResponse struct {

	// 应用实例列表
	AppInstances   *[]QueryAppInstanceResp `json:"app_instances,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListAppInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListAppInstancesResponse", string(data)}, " ")
}
