package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePipeResponse Response Object
type DeletePipeResponse struct {

	// UUID
	PipeId *string `json:"pipe_id,omitempty"`

	// 毫秒时间戳
	DeleteTime     *int64 `json:"delete_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DeletePipeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePipeResponse struct{}"
	}

	return strings.Join([]string{"DeletePipeResponse", string(data)}, " ")
}
