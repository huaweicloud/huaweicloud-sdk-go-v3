package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunSemanticParserRequest struct {
	Body *IntentReq `json:"body,omitempty"`
}

func (o RunSemanticParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSemanticParserRequest struct{}"
	}

	return strings.Join([]string{"RunSemanticParserRequest", string(data)}, " ")
}
