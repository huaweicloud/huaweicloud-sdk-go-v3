package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type SwitchConfigurationResponse struct {

	// 应用参数模板的异步任务ID。
	JobId          *string `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchConfigurationResponse struct{}"
	}

	return strings.Join([]string{"SwitchConfigurationResponse", string(data)}, " ")
}
