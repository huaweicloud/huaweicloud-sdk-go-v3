package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlinkTemplateResponse Response Object
type DeleteFlinkTemplateResponse struct {

	// 响应正确与否的标志，true表示成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 消息内容。
	Message *string `json:"message,omitempty"`

	Template       *DeleteTemplateRespTemplate `json:"template,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o DeleteFlinkTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteFlinkTemplateResponse", string(data)}, " ")
}
