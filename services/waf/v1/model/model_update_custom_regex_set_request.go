package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomRegexSetRequest Request Object
type UpdateCustomRegexSetRequest struct {

	// 正则库ID
	RegexSetId string `json:"regex_set_id"`

	Body *AiguardCustomRegexSetUpdateBody `json:"body,omitempty"`
}

func (o UpdateCustomRegexSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomRegexSetRequest struct{}"
	}

	return strings.Join([]string{"UpdateCustomRegexSetRequest", string(data)}, " ")
}
