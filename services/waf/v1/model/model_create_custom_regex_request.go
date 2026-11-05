package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomRegexRequest Request Object
type CreateCustomRegexRequest struct {

	// 正则库ID
	RegexSetId string `json:"regex_set_id"`

	Body *AiguardCustomRegexBody `json:"body,omitempty"`
}

func (o CreateCustomRegexRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomRegexRequest struct{}"
	}

	return strings.Join([]string{"CreateCustomRegexRequest", string(data)}, " ")
}
