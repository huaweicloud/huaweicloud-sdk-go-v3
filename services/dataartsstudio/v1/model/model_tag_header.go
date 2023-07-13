package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagHeader struct {

	// 资产名称
	Name *string `json:"name,omitempty"`

	// 标签描述
	Dexcription *interface{} `json:"dexcription,omitempty"`

	// 标签的名称
	DisplayText *string `json:"display_text,omitempty"`

	// 关联的guid
	RelationGuid *string `json:"relation_guid,omitempty"`

	// 标签关联的guid
	TagGuid *string `json:"tag_guid,omitempty"`
}

func (o TagHeader) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagHeader struct{}"
	}

	return strings.Join([]string{"TagHeader", string(data)}, " ")
}
