package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InvokeGaussMySqlInstanceSwitchOverResponse Response Object
type InvokeGaussMySqlInstanceSwitchOverResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o InvokeGaussMySqlInstanceSwitchOverResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InvokeGaussMySqlInstanceSwitchOverResponse struct{}"
	}

	return strings.Join([]string{"InvokeGaussMySqlInstanceSwitchOverResponse", string(data)}, " ")
}
