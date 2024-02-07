package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetProfileImageRequest Request Object
type SetProfileImageRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	Body *SetProfileImageRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o SetProfileImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetProfileImageRequest struct{}"
	}

	return strings.Join([]string{"SetProfileImageRequest", string(data)}, " ")
}
