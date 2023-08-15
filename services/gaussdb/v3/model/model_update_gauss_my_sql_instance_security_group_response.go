package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlInstanceSecurityGroupResponse Response Object
type UpdateGaussMySqlInstanceSecurityGroupResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlInstanceSecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceSecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceSecurityGroupResponse", string(data)}, " ")
}
