package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserQuotasRequest Request Object
type UpdateUserQuotasRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 子账户（用户）ID。
	UserId string `json:"user_id"`

	Body *UpdateUserQuotaInfo `json:"body,omitempty"`
}

func (o UpdateUserQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserQuotasRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserQuotasRequest", string(data)}, " ")
}
