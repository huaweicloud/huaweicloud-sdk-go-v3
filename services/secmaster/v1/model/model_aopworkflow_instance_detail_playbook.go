package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AopworkflowInstanceDetailPlaybook 剧本对象
type AopworkflowInstanceDetailPlaybook struct {

	// **参数解释**: 剧本英文名字 **约束限制**: 不涉及
	EnName *string `json:"en_name,omitempty"`

	// **参数解释**: 剧本的ID **约束限制**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 剧本的名字 **约束限制**: 不涉及
	Name *string `json:"name,omitempty"`
}

func (o AopworkflowInstanceDetailPlaybook) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopworkflowInstanceDetailPlaybook struct{}"
	}

	return strings.Join([]string{"AopworkflowInstanceDetailPlaybook", string(data)}, " ")
}
