package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTranscodeTemplateResponse Response Object
type DeleteTranscodeTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTranscodeTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTranscodeTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteTranscodeTemplateResponse", string(data)}, " ")
}
