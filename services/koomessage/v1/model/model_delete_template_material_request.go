package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateMaterialRequest Request Object
type DeleteTemplateMaterialRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *DeleteTemplateMaterialRequestBody `json:"body,omitempty"`
}

func (o DeleteTemplateMaterialRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateMaterialRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateMaterialRequest", string(data)}, " ")
}
