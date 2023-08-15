package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateRequest Request Object
type ShowTemplateRequest struct {

	// 模板ID
	Id string `json:"id"`
}

func (o ShowTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateRequest", string(data)}, " ")
}
