package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApplicableInstancesResponse Response Object
type ShowApplicableInstancesResponse struct {

	// 实例列表
	Instances *[]ApplicableInstanceRsp `json:"instances,omitempty"`

	// 应用参数的实例数量限制。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowApplicableInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplicableInstancesResponse struct{}"
	}

	return strings.Join([]string{"ShowApplicableInstancesResponse", string(data)}, " ")
}
