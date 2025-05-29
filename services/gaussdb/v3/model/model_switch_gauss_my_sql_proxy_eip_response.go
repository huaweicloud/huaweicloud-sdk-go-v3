package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlProxyEipResponse Response Object
type SwitchGaussMySqlProxyEipResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchGaussMySqlProxyEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlProxyEipResponse struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlProxyEipResponse", string(data)}, " ")
}
