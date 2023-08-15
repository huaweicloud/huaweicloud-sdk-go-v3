package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGaussMySqlInstanceInternalIpResponse Response Object
type UpdateGaussMySqlInstanceInternalIpResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlInstanceInternalIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceInternalIpResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceInternalIpResponse", string(data)}, " ")
}
