package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StructTemplateModel 更新或者查询结构化模板对象
type StructTemplateModel struct {
}

func (o StructTemplateModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructTemplateModel struct{}"
	}

	return strings.Join([]string{"StructTemplateModel", string(data)}, " ")
}
