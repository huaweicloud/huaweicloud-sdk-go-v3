package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllScriptByNameRequest Request Object
type ListAllScriptByNameRequest struct {
	Body *SearchScriptsRequestBody `json:"body,omitempty"`
}

func (o ListAllScriptByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllScriptByNameRequest struct{}"
	}

	return strings.Join([]string{"ListAllScriptByNameRequest", string(data)}, " ")
}
