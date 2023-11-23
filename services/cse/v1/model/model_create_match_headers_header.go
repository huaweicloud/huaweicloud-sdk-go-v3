package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMatchHeadersHeader 匹配Header的规则。
type CreateMatchHeadersHeader struct {

	// 精确匹配值。
	Exact *string `json:"exact,omitempty"`

	// 是否区分大小写。
	CaseInsensitive *bool `json:"caseInsensitive,omitempty"`
}

func (o CreateMatchHeadersHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMatchHeadersHeader struct{}"
	}

	return strings.Join([]string{"CreateMatchHeadersHeader", string(data)}, " ")
}
