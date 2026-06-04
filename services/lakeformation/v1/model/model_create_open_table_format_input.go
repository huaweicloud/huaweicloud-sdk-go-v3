package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOpenTableFormatInput 开源湖表格式结构体
type CreateOpenTableFormatInput struct {
	CreateIcebergTableInput *CreateIcebergTableInput `json:"create_iceberg_table_input,omitempty"`

	CreateLanceTableInput *CreateLanceTableInput `json:"create_lance_table_input,omitempty"`
}

func (o CreateOpenTableFormatInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOpenTableFormatInput struct{}"
	}

	return strings.Join([]string{"CreateOpenTableFormatInput", string(data)}, " ")
}
