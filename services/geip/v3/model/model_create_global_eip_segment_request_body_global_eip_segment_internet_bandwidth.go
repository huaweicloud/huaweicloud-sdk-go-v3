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

	// - 功能说明：全域公网带宽名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 全域弹性公网IP段标签
	Tags *[]CreateGlobalEipRequestBodyGlobalEipTags `json:"tags,omitempty"`

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
