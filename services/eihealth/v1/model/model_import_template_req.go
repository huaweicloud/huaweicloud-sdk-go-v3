package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 导入模板请求体
type ImportTemplateReq struct {

	// 源项目id
	SourceProjectId string `json:"source_project_id"`

	// 导入模板列表
	ImportTemplates []TemplateSrcReq `json:"import_templates"`
}

func (o ImportTemplateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportTemplateReq struct{}"
	}

	return strings.Join([]string{"ImportTemplateReq", string(data)}, " ")
}
