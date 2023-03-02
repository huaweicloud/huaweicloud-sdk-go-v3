package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiVersionsLinks struct {

	// href属性
	Href *string `json:"href,omitempty"`

	// rel属性
	Rel *string `json:"rel,omitempty"`
}

func (o ApiVersionsLinks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersionsLinks struct{}"
	}

	return strings.Join([]string{"ApiVersionsLinks", string(data)}, " ")
}
