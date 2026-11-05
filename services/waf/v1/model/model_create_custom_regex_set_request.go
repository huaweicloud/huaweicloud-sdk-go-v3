package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomRegexSetRequest Request Object
type CreateCustomRegexSetRequest struct {
	Body *AiguardCustomRegexSetBody `json:"body,omitempty"`
}

func (o CreateCustomRegexSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomRegexSetRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomRegexSetRequest", string(data)}, " ")
}
