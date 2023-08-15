package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobResponse Response Object
type UpdateJobResponse struct {

	// 错误码。
	ErrorCode string `json:"error_code"`

	// 错误描述。
	ErrorMsg string `json:"error_msg"`

	// 任务ID。
	Id string `json:"id"`

	// 任务名称。
	Name string `json:"name"`

	// 操作结果。
	Status         string `json:"status"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobResponse", string(data)}, " ")
}
