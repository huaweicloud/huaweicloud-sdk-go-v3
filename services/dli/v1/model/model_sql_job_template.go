package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SqlJobTemplate struct {

	// SQL模板ID。
	SqlId *string `json:"sql_id,omitempty"`

	// SQL模板名称。
	SqlName *string `json:"sql_name,omitempty"`

	// SQL模板文本。
	Sql *string `json:"sql,omitempty"`

	// SQL模板描述信息。
	Description *string `json:"description,omitempty"`

	// 分组名称。
	Group *string `json:"group,omitempty"`

	// SQL模板的创建者。
	Owner *string `json:"owner,omitempty"`
}

func (o SqlJobTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlJobTemplate struct{}"
	}

	return strings.Join([]string{"SqlJobTemplate", string(data)}, " ")
}
