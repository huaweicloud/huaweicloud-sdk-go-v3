package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAppIdRequest Request Object
type AddAppIdRequest struct {

	// 请求requestId，用来标识一路请求，用于问题跟踪定位，建议使用UUID，若不携带，则后台自动生成。
	XRequestId *string `json:"X-Request-Id,omitempty"`

	Body *AddAppIdRequestBody `json:"body,omitempty"`
}

func (o AddAppIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAppIdRequest struct{}"
	}

	return strings.Join([]string{"AddAppIdRequest", string(data)}, " ")
}
