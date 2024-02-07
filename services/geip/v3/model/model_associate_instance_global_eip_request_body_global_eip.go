package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateInstanceGlobalEipRequestBodyGlobalEip 绑定后端实例请求体对象
type AssociateInstanceGlobalEipRequestBodyGlobalEip struct {
	AssociateInstanceInfo *AssociateInstanceGlobalEipRequestBodyGlobalEipAssociateInstanceInfo `json:"associate_instance_info,omitempty"`

	GcBandwidthInfo *AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfo `json:"gc_bandwidth_info,omitempty"`
}

func (o AssociateInstanceGlobalEipRequestBodyGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateInstanceGlobalEipRequestBodyGlobalEip struct{}"
	}

	return strings.Join([]string{"AssociateInstanceGlobalEipRequestBodyGlobalEip", string(data)}, " ")
}
