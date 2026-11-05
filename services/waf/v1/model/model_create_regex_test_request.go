package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRegexTestRequest Request Object
type CreateRegexTestRequest struct {
	Body *RegexTestRequest `json:"body,omitempty"`
}

func (o CreateRegexTestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRegexTestRequest struct{}"
	}

	return strings.Join([]string{"CreateRegexTestRequest", string(data)}, " ")
}
