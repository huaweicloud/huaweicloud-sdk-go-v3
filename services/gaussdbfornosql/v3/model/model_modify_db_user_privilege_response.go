package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ModifyDbUserPrivilegeResponse struct {

	// 任务ID.
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyDbUserPrivilegeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDbUserPrivilegeResponse struct{}"
	}

	return strings.Join([]string{"ModifyDbUserPrivilegeResponse", string(data)}, " ")
}
