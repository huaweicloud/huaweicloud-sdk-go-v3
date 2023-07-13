package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchemaInfo 源库中的schema信息。
type SchemaInfo struct {

	// 是否选择全部schema。
	IsSelectAllSchemas bool `json:"is_select_all_schemas"`

	// 需要评估的源库schema列表。is_select_all_schemas为false时，必填。
	SchemasList *[]string `json:"schemas_list,omitempty"`
}

func (o SchemaInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchemaInfo struct{}"
	}

	return strings.Join([]string{"SchemaInfo", string(data)}, " ")
}
