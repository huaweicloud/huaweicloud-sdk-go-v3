package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateRecordCallbackRequest struct {
	// 使用AK/SK方式认证时必选，携带的鉴权信息。

	Authorization *string `json:"Authorization,omitempty"`
	// 使用AK/SK方式认证时必选，请求的发生时间。

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`
	// 使用AK/SK方式认证时必选，携带项目ID信息。

	XProjectId *string `json:"X-Project-Id,omitempty"`
	// 应用id

	AppId string `json:"app_id"`

	Body *AppCallbackUrlReq `json:"body,omitempty"`
}

func (o UpdateRecordCallbackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRecordCallbackRequest struct{}"
	}

	return strings.Join([]string{"UpdateRecordCallbackRequest", string(data)}, " ")
}
