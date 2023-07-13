package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchemasList schema信息
type SchemasList struct {

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`

	// schema描述
	Description *string `json:"description,omitempty"`
}

func (o SchemasList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemasList struct{}"
	}

	return strings.Join([]string{"SchemasList", string(data)}, " ")
}
