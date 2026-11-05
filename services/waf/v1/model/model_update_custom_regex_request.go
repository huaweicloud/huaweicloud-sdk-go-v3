package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomRegexRequest Request Object
type UpdateCustomRegexRequest struct {

	// 正则库ID
	RegexSetId string `json:"regex_set_id"`

	// 正则ID
	RegexId string `json:"regex_id"`

	Body *AiguardCustomRegexUpdateBody `json:"body,omitempty"`
}

func (o UpdateCustomRegexRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomRegexRequest struct{}"
	}

	return strings.Join([]string{"UpdateCustomRegexRequest", string(data)}, " ")
}
