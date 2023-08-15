package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFlinkTemplateRequest Request Object
type DeleteFlinkTemplateRequest struct {

	// 模板ID。
	TemplateId int64 `json:"template_id"`
}

func (o DeleteFlinkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFlinkTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteFlinkTemplateRequest", string(data)}, " ")
}
