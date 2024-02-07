package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetUserProfileImageRequest Request Object
type SetUserProfileImageRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	// 设置头像的企业用户id
	UserId string `json:"user_id"`

	Body *SetUserProfileImageRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o SetUserProfileImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetUserProfileImageRequest struct{}"
	}

	return strings.Join([]string{"SetUserProfileImageRequest", string(data)}, " ")
}
