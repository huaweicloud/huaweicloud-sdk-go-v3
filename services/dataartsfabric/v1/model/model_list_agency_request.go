package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgencyRequest Request Object
type ListAgencyRequest struct {

	// 策略类型。支持模糊匹配
	PolicyType *string `json:"policy_type,omitempty"`
}

func (o ListAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgencyRequest struct{}"
	}

	return strings.Join([]string{"ListAgencyRequest", string(data)}, " ")
}
