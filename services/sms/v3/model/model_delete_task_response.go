package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaskResponse Response Object
type DeleteTaskResponse struct {

	// 删除指定ID的迁移任务成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaskResponse struct{}"
	}

	return strings.Join([]string{"DeleteTaskResponse", string(data)}, " ")
}
