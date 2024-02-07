package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGlobalEipRequestBodyGlobalEip 请求体信息
type CreateGlobalEipRequestBodyGlobalEip struct {

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`

	// 全域弹性公网IP池子名称
	GeipPoolName string `json:"geip_pool_name"`

	// 接入点信息
	AccessSite string `json:"access_site"`

	// 全域公网带宽的ID
	InternetBandwidthId *string `json:"internet_bandwidth_id,omitempty"`

	// 全域弹性公网IP标签
	Tags *[]CreateGlobalEipRequestBodyGlobalEipTags `json:"tags,omitempty"`

	// 资源的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateGlobalEipRequestBodyGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGlobalEipRequestBodyGlobalEip struct{}"
	}

	return strings.Join([]string{"CreateGlobalEipRequestBodyGlobalEip", string(data)}, " ")
}
