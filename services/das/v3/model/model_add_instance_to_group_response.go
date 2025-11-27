package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddInstanceToGroupResponse Response Object
type AddInstanceToGroupResponse struct {

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 成功的实例ID列表
	SuccessInstanceIds *[]string `json:"success_instance_ids,omitempty"`

	// 失败的实例ID列表
	FailedInstanceIds *[]string `json:"failed_instance_ids,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o AddInstanceToGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddInstanceToGroupResponse struct{}"
	}

	return strings.Join([]string{"AddInstanceToGroupResponse", string(data)}, " ")
}
