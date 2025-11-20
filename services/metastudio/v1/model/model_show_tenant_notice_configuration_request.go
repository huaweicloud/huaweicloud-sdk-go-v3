package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantNoticeConfigurationRequest Request Object
type ShowTenantNoticeConfigurationRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 通知类型。 * RESOURCE_EXPIRE：资源过期通知 * RESOURCE_LEFT：资源剩余量通知
	Type string `json:"type"`
}

func (o ShowTenantNoticeConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantNoticeConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ShowTenantNoticeConfigurationRequest", string(data)}, " ")
}
