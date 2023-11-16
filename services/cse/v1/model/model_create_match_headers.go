package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMatchHeaders headers
type CreateMatchHeaders struct {
	Aaaa *CreateMatchHeadersAaaa `json:"aaaa,omitempty"`
}

func (o CreateMatchHeaders) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMatchHeaders struct{}"
	}

	return strings.Join([]string{"CreateMatchHeaders", string(data)}, " ")
}
