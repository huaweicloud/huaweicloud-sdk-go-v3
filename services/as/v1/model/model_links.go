package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Links struct {

	// API的URL地址。
	Href *string `json:"href,omitempty" xml:"href"`

	// API的URL依赖。
	Rel *string `json:"rel,omitempty" xml:"rel"`
}

func (o Links) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Links struct{}"
	}

	return strings.Join([]string{"Links", string(data)}, " ")
}
