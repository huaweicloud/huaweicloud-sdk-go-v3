package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeCustomTemplateResponse Response Object
type RecognizeCustomTemplateResponse struct {

	// 调用成功时表示调用结果。 调用失败时无此字段。
	Result *interface{} `json:"result,omitempty"`

	// 调用成功时返回调用模板id。 调用失败时无此字段。
	TemplateId *string `json:"template_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeCustomTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeCustomTemplateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeCustomTemplateResponse", string(data)}, " ")
}
