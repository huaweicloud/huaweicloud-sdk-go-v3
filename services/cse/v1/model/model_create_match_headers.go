package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMatchHeaders 匹配的Headers。
type CreateMatchHeaders struct {
	Header *CreateMatchHeadersHeader `json:"&lt;header&gt;,omitempty"`
}

func (o CreateMatchHeaders) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMatchHeaders struct{}"
	}

	return strings.Join([]string{"CreateMatchHeaders", string(data)}, " ")
}
