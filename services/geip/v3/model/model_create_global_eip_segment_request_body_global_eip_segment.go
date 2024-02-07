package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalEipSegmentRequestBodyGlobalEipSegment 创建全域弹性公网IP段请求对象信息
type CreateGlobalEipSegmentRequestBodyGlobalEipSegment struct {

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`

	// 全域弹性公网IP池子名称
	GeipPoolName string `json:"geip_pool_name"`

	// 接入点信息
	AccessSite string `json:"access_site"`

	// 掩码长度。取值范围由GET /v3/{domain_id}/global-eip-segments/support-masks接口提供
	Mask int32 `json:"mask"`

	InternetBandwidth *CreateGlobalEipSegmentRequestBodyGlobalEipSegmentInternetBandwidth `json:"internet_bandwidth,omitempty"`

	// 全域弹性公网IP段标签
	Tags *[]CreateGlobalEipRequestBodyGlobalEipTags `json:"tags,omitempty"`

	// 资源的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateGlobalEipSegmentRequestBodyGlobalEipSegment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEipSegmentRequestBodyGlobalEipSegment struct{}"
	}

	return strings.Join([]string{"CreateGlobalEipSegmentRequestBodyGlobalEipSegment", string(data)}, " ")
}
