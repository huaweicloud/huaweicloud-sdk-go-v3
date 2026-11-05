package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomRegexRequest Request Object
type DeleteCustomRegexRequest struct {

	// 正则库ID
	RegexSetId string `json:"regex_set_id"`

	// 正则ID
	RegexId string `json:"regex_id"`
}

func (o DeleteCustomRegexRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomRegexRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomRegexRequest", string(data)}, " ")
}
