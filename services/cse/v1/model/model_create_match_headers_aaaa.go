package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMatchHeadersAaaa aaaa
type CreateMatchHeadersAaaa struct {

	// exact
	Exact *string `json:"exact,omitempty"`

	// caseInsensitive
	CaseInsensitive *bool `json:"caseInsensitive,omitempty"`
}

func (o CreateMatchHeadersAaaa) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMatchHeadersAaaa struct{}"
	}

	return strings.Join([]string{"CreateMatchHeadersAaaa", string(data)}, " ")
}
