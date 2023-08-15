package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVideoMotionCaptureJobRequest Request Object
type CreateVideoMotionCaptureJobRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 开发者应用作为资产权属的可选字段。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 测试用户判断。
	XUserPrivilege *string `json:"X-User-Privilege,omitempty"`

	Body *VideoMotionCaptureJobReq `json:"body,omitempty"`
}

func (o CreateVideoMotionCaptureJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVideoMotionCaptureJobRequest struct{}"
	}

	return strings.Join([]string{"CreateVideoMotionCaptureJobRequest", string(data)}, " ")
}
