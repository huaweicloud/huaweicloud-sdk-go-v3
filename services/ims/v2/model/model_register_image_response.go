package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RegisterImageResponse struct {
	// 异步任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageResponse struct{}"
	}

	return strings.Join([]string{"RegisterImageResponse", string(data)}, " ")
}
