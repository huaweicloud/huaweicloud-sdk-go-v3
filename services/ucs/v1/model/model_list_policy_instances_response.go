package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyInstancesResponse Response Object
type ListPolicyInstancesResponse struct {

	// 策略实例列表
	Items *[]UcsConstraint `json:"items,omitempty"`

	// API类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion     *string `json:"apiVersion,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPolicyInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyInstancesResponse", string(data)}, " ")
}
