package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAimTemplateReportsRequest Request Object
type ListAimTemplateReportsRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *ListAimTemplateReportsRequestBody `json:"body,omitempty"`
}

func (o ListAimTemplateReportsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAimTemplateReportsRequest struct{}"
	}

	return strings.Join([]string{"ListAimTemplateReportsRequest", string(data)}, " ")
}
