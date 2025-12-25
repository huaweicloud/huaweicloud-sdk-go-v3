package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePreProcessRulesResponseBodyData 数据
type CreatePreProcessRulesResponseBodyData struct {

	// 映射id
	ProjectId *string `json:"project_id,omitempty"`

	// 映射id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 映射id
	MappingId *string `json:"mapping_id,omitempty"`

	// 预处理规则列表
	PreprocessRules *[]CreatePreProcessRulesResponseBodyDataPreprocessRules `json:"preprocess_rules,omitempty"`
}

func (o CreatePreProcessRulesResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePreProcessRulesResponseBodyData struct{}"
	}

	return strings.Join([]string{"CreatePreProcessRulesResponseBodyData", string(data)}, " ")
}
