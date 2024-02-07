package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInternetBandwidthRequestBodyInternetBandwidth 全域公网带宽信息
type CreateInternetBandwidthRequestBodyInternetBandwidth struct {

	// 全域公网带宽大小（入云方向）
	IngressSize *int32 `json:"ingress_size,omitempty"`

	// 计费模式
	ChargeMode string `json:"charge_mode"`

	// 线路
	Isp string `json:"isp"`

	// 接入点信息
	AccessSite string `json:"access_site"`

	// 全域公网带宽大小（出云方向）
	Size int32 `json:"size"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`

	// 全域公网带宽标签
	Tags *[]CreateGlobalEipRequestBodyGlobalEipTags `json:"tags,omitempty"`

	// 资源的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 全域公网带宽类型
	Type *string `json:"type,omitempty"`
}

func (o CreateInternetBandwidthRequestBodyInternetBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInternetBandwidthRequestBodyInternetBandwidth struct{}"
	}

	return strings.Join([]string{"CreateInternetBandwidthRequestBodyInternetBandwidth", string(data)}, " ")
}
