package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomRegexSetRequest Request Object
type ShowCustomRegexSetRequest struct {

	// 正则库ID
	RegexSetId string `json:"regex_set_id"`
}

func (o ShowCustomRegexSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomRegexSetRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomRegexSetRequest", string(data)}, " ")
}
