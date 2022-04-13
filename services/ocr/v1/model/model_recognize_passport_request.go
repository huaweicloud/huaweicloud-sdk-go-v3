package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RecognizePassportRequest struct {
	Body *PassportRequestBody `json:"body,omitempty"`
}

func (o RecognizePassportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizePassportRequest struct{}"
	}

	return strings.Join([]string{"RecognizePassportRequest", string(data)}, " ")
}
