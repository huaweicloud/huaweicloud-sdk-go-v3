package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeUserPrivilegeGroupResponse Response Object
type ChangeUserPrivilegeGroupResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeUserPrivilegeGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeUserPrivilegeGroupResponse struct{}"
	}

	return strings.Join([]string{"ChangeUserPrivilegeGroupResponse", string(data)}, " ")
}
