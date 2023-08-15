package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateRequest Request Object
type DeleteTemplateRequest struct {

	// 模板ID
	Id string `json:"id"`
}

func (o DeleteTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteTemplateRequest", string(data)}, " ")
}
