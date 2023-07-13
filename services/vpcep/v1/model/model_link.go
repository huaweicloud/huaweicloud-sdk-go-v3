package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Link API的url地址
type Link struct {

	// 当前API版本的引用地址。
	Href *string `json:"href,omitempty"`

	// 发送的实体的MIME类型，取值为application/json。
	Type *string `json:"type,omitempty"`

	// 当前API版本和被引用地址的关系。
	Rel *string `json:"rel,omitempty"`
}

func (o Link) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Link struct{}"
	}

	return strings.Join([]string{"Link", string(data)}, " ")
}
