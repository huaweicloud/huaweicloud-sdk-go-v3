package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchGaussMySqlConfigurationResponse Response Object
type SwitchGaussMySqlConfigurationResponse struct {

	// 应用参数模板的任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchGaussMySqlConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchGaussMySqlConfigurationResponse struct{}"
	}

	return strings.Join([]string{"SwitchGaussMySqlConfigurationResponse", string(data)}, " ")
}
