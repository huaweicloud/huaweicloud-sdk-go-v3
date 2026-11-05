package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FailedCustomRegexSet struct {

	// 失败的正则库ID
	RegexSetId *string `json:"regex_set_id,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`
}

func (o FailedCustomRegexSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FailedCustomRegexSet struct{}"
	}

	return strings.Join([]string{"FailedCustomRegexSet", string(data)}, " ")
}
