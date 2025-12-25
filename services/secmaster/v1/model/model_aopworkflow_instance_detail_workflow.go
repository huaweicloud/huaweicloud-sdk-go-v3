package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AopworkflowInstanceDetailWorkflow 流程对象
type AopworkflowInstanceDetailWorkflow struct {

	// **参数解释**: 流程的ID **约束限制**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 流程的中文名字 **约束限制**: 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**: 流程的英文名字 **约束限制**: 不涉及
	NameEn *string `json:"name_en,omitempty"`

	// **参数解释**: 流程的版本 **约束限制**: 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**: 流程的版本ID **约束限制**: 不涉及
	VersionId *string `json:"version_id,omitempty"`
}

func (o AopworkflowInstanceDetailWorkflow) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopworkflowInstanceDetailWorkflow struct{}"
	}

	return strings.Join([]string{"AopworkflowInstanceDetailWorkflow", string(data)}, " ")
}
