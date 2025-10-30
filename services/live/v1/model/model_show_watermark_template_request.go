package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWatermarkTemplateRequest Request Object
type ShowWatermarkTemplateRequest struct {

	// 模板ID
	Id string `json:"id"`
}

func (o ShowWatermarkTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowWatermarkTemplateRequest", string(data)}, " ")
}
