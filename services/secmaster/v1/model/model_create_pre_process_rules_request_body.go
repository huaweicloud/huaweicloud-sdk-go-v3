package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePreProcessRulesRequestBody struct {

	// 映射id
	ProjectId *string `json:"project_id,omitempty"`

	// 映射id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 映射id
	MappingId string `json:"mapping_id"`

	// 预处理规则列表
	PreprocessRules *[]CreatePreProcessRulesRequestBodyPreprocessRules `json:"preprocess_rules,omitempty"`
}

func (o CreatePreProcessRulesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePreProcessRulesRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePreProcessRulesRequestBody", string(data)}, " ")
}
