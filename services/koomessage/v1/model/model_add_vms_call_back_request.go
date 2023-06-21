package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddVmsCallBackRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *AddVmsCallBackRequestBody `json:"body,omitempty"`
}

func (o AddVmsCallBackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVmsCallBackRequest struct{}"
	}

	return strings.Join([]string{"AddVmsCallBackRequest", string(data)}, " ")
}
