package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskInstancesResponse Response Object
type ShowTaskInstancesResponse struct {

	// 实例个数
	Count *int32 `json:"count,omitempty"`

	// 实例响应结构体
	Instances      *[]TaskInstanceRsp `json:"instances,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowTaskInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskInstancesResponse", string(data)}, " ")
}
