package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletedPolicy 待删除的安全组策略
type DeletedPolicy struct {

	// 策略ID
	PolicyId string `json:"policy_id"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o DeletedPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletedPolicy struct{}"
	}

	return strings.Join([]string{"DeletedPolicy", string(data)}, " ")
}
