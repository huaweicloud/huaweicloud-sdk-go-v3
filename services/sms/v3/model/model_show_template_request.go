package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateRequest Request Object
type ShowTemplateRequest struct {

	// 需要查询的模板信息的ID
	Id string `json:"id"`
}

func (o ShowTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateRequest", string(data)}, " ")
}
