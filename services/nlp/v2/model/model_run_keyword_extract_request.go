package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunKeywordExtractRequest Request Object
type RunKeywordExtractRequest struct {
	Body *KeywordExtractReq `json:"body,omitempty"`
}

func (o RunKeywordExtractRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunKeywordExtractRequest struct{}"
	}

	return strings.Join([]string{"RunKeywordExtractRequest", string(data)}, " ")
}
