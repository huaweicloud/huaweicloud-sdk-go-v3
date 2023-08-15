package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTranscodeTemplateResponse Response Object
type UpdateTranscodeTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTranscodeTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTranscodeTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateTranscodeTemplateResponse", string(data)}, " ")
}
