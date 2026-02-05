package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateSpeechRequest Request Object
type GenerateSpeechRequest struct {
	Body *GenerateSpeechRequestBody `json:"body,omitempty"`
}

func (o GenerateSpeechRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateSpeechRequest struct{}"
	}

	return strings.Join([]string{"GenerateSpeechRequest", string(data)}, " ")
}
