package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlProxyAltResponse Response Object
type SwitchGaussMySqlProxyAltResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchGaussMySqlProxyAltResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxyAltResponse struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxyAltResponse", string(data)}, " ")
}
