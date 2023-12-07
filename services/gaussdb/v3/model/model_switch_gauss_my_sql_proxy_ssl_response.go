package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlProxySslResponse Response Object
type SwitchGaussMySqlProxySslResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchGaussMySqlProxySslResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxySslResponse struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxySslResponse", string(data)}, " ")
}
