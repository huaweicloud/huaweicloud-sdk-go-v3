package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyGroupForTemplate 策略组模板。
type PolicyGroupForTemplate struct {

	// 策略组名称。
	PolicyGroupName string `json:"policy_group_name"`

	// 策略组描述。
	Description *string `json:"description,omitempty"`

	// 策略。
	Policies string `json:"policies"`

	// 策略组类型。
	PolicyGroupType *string `json:"policy_group_type,omitempty"`
}

func (o PolicyGroupForTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyGroupForTemplate struct{}"
	}

	return strings.Join([]string{"PolicyGroupForTemplate", string(data)}, " ")
}
