package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTranscodeTemplateRequest Request Object
type DeleteTranscodeTemplateRequest struct {

	// 模板id
	GroupId string `json:"group_id"`
}

func (o DeleteTranscodeTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTranscodeTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteTranscodeTemplateRequest", string(data)}, " ")
}
