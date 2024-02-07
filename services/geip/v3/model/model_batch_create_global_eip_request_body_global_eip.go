package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGlobalEipRequestBodyGlobalEip 批量创建全域弹性公网IP请求体对象
type BatchCreateGlobalEipRequestBodyGlobalEip struct {

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 用户自定义的资源描述
	Description *string `json:"description,omitempty"`

	// 全域弹性公网IP池子名称
	GeipPoolName string `json:"geip_pool_name"`

	// 接入点信息
	AccessSite string `json:"access_site"`

	InternetBandwidthInfo *BatchCreateGlobalEipRequestBodyGlobalEipInternetBandwidthInfo `json:"internet_bandwidth_info"`

	// 批创个数
	Count *int32 `json:"count,omitempty"`

	// 全域弹性公网IP标签
	Tags *[]CreateGlobalEipRequestBodyGlobalEipTags `json:"tags,omitempty"`

	// 资源的企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o BatchCreateGlobalEipRequestBodyGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGlobalEipRequestBodyGlobalEip struct{}"
	}

	return strings.Join([]string{"BatchCreateGlobalEipRequestBodyGlobalEip", string(data)}, " ")
}
