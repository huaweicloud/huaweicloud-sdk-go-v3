package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunSemanticParserRequest struct {
	Body *IntentReq `json:"body,omitempty" xml:"body"`
}

func (o RunSemanticParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSemanticParserRequest struct{}"
	}

	return strings.Join([]string{"RunSemanticParserRequest", string(data)}, " ")
}
