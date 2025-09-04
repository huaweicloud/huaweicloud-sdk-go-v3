package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateRequest Request Object
type ShowTemplateRequest struct {

	// uuid
	Uuid string `json:"uuid"`
}

func (o ShowTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateRequest", string(data)}, " ")
}
