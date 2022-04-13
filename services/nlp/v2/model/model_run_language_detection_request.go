package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunLanguageDetectionRequest struct {
	Body *LanguageDetectionReq `json:"body,omitempty"`
}

func (o RunLanguageDetectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunLanguageDetectionRequest struct{}"
	}

	return strings.Join([]string{"RunLanguageDetectionRequest", string(data)}, " ")
}
