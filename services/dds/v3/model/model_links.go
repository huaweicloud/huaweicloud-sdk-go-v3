package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 对应API的链接信息，v3版本该字段为[]。
type Links struct {

	// 对应该API的URL，该字段为\"\"。
	Href *string `json:"href,omitempty"`

	// 值为“self”，表示URL为本地链接。
	Rel *string `json:"rel,omitempty"`
}

func (o Links) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Links struct{}"
	}

	return strings.Join([]string{"Links", string(data)}, " ")
}
