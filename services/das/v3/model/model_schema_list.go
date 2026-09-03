package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchemaList Schema信息
type SchemaList struct {

	// schema名称
	SchemaName *string `json:"schema_name,omitempty"`
}

func (o SchemaList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaList struct{}"
	}

	return strings.Join([]string{"SchemaList", string(data)}, " ")
}
