package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAnalyzerRequest Request Object
type CreateAnalyzerRequest struct {
	Body *CreateAnalyzerReqBody `json:"body,omitempty"`
}

func (o CreateAnalyzerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAnalyzerRequest struct{}"
	}

	return strings.Join([]string{"CreateAnalyzerRequest", string(data)}, " ")
}
