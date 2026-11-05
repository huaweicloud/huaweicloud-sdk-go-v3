package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomRegexSetRequest Request Object
type DeleteCustomRegexSetRequest struct {

	// 正则库ID
	RegexSetId string `json:"regex_set_id"`
}

func (o DeleteCustomRegexSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomRegexSetRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomRegexSetRequest", string(data)}, " ")
}
