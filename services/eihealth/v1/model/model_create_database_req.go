package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建实例请求体
type CreateDatabaseReq struct {

	// 实例名称
	Name string `json:"name"`

	// 模板id
	TemplateId string `json:"template_id"`

	// 实例描述
	Description *string `json:"description,omitempty"`

	ImportData *ImportDatabaseDataReq `json:"import_data,omitempty"`
}

func (o CreateDatabaseReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseReq struct{}"
	}

	return strings.Join([]string{"CreateDatabaseReq", string(data)}, " ")
}
