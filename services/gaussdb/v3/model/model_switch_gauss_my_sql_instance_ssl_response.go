package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlInstanceSslResponse Response Object
type SwitchGaussMySqlInstanceSslResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchGaussMySqlInstanceSslResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlInstanceSslResponse struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlInstanceSslResponse", string(data)}, " ")
}
