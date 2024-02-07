package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGlobalEipSegmentRequestBodyGlobalEipSegmentInternetBandwidth struct {

	// 全域公网带宽的ID
	Id *string `json:"id,omitempty"`

	// 全域公网带宽大小（入云方向）
	IngressSize *int32 `json:"ingress_size,omitempty"`

	// 计费模式
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 全域公网带宽大小（出云方向）
	Size *int32 `json:"size,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 全域弹性公网IP段标签
	Tags *[]AssociateInstanceGlobalEipRequestBodyGlobalEipGcBandwidthInfoTags `json:"tags,omitempty"`

	// 全域公网带宽类型
	Type *string `json:"type,omitempty"`
}

func (o CreateGlobalEipSegmentRequestBodyGlobalEipSegmentInternetBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEipSegmentRequestBodyGlobalEipSegmentInternetBandwidth struct{}"
	}

	return strings.Join([]string{"CreateGlobalEipSegmentRequestBodyGlobalEipSegmentInternetBandwidth", string(data)}, " ")
}
