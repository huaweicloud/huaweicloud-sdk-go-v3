package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ApplyConfigurationResponse struct {
	// 应用参数模板的异步任务ID。

	JobId *string `json:"job_id,omitempty"`
	// 应用参数模板任务是否提交成功。 - 取值为“true”，表示任务提交成功。 - 取值为“false”，表示任务提交失败。

	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ApplyConfigurationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApplyConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationResponse", string(data)}, " ")
}
