package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建评估项目请求体。
type CreateEvaluationProjectReq struct {

	// 评估项目名称。长度为5-50个字符，以英文字母开头，英文字母或数字结束，允许包含下划线和中划线。不允许重复。
	EvaluationProjectName string `json:"evaluation_project_name"`

	SourceDbInfo *SourceDbInfo `json:"source_db_info"`

	SchemasInfo *SchemaInfo `json:"schemas_info"`

	ObjectsTypeInfo *ObjectTypeInfo `json:"objects_type_info"`
}

func (o CreateEvaluationProjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEvaluationProjectReq struct{}"
	}

	return strings.Join([]string{"CreateEvaluationProjectReq", string(data)}, " ")
}
