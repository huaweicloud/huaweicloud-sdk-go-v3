package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建模板请求体
type CreateTemplateReq struct {

	// 模板名称
	Name string `json:"name"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 数据库列信息
	Columns []DatabaseColumnDto `json:"columns"`
}

func (o CreateTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplateReq struct{}"
	}

	return strings.Join([]string{"CreateTemplateReq", string(data)}, " ")
}
