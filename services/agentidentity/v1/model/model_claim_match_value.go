package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClaimMatchValue The value or values to match for.
type ClaimMatchValue struct {

	// The string value to match for.
	MatchValueString *string `json:"match_value_string,omitempty"`

	MatchValueStringList *[]string `json:"match_value_string_list,omitempty"`
}

func (o ClaimMatchValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClaimMatchValue struct{}"
	}

	return strings.Join([]string{"ClaimMatchValue", string(data)}, " ")
}
