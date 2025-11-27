package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyJobsResponse Response Object
type ListPolicyJobsResponse struct {

	// Job列表
	Items *[]UcsJob `json:"items,omitempty"`

	// API类型
	Kind *string `json:"kind,omitempty"`

	// API版本
	ApiVersion     *string `json:"apiVersion,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPolicyJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyJobsResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyJobsResponse", string(data)}, " ")
}
