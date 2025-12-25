package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AopworkflowInstanceDetailDataclass 数据类对象
type AopworkflowInstanceDetailDataclass struct {

	// **参数解释**: 数据类的英文名字 **约束限制**: 不涉及
	EnName *string `json:"en_name,omitempty"`

	// **参数解释**: 数据类的ID **约束限制**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 数据类的中文名字 **约束限制**: 不涉及
	Name *string `json:"name,omitempty"`
}

func (o AopworkflowInstanceDetailDataclass) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopworkflowInstanceDetailDataclass struct{}"
	}

	return strings.Join([]string{"AopworkflowInstanceDetailDataclass", string(data)}, " ")
}
