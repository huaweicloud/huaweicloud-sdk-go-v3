package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartEncodeServerResponse Response Object
type RestartEncodeServerResponse struct {

	// 请求的唯一标识ID
	RequestId *string `json:"request_id,omitempty"`

	// 任务信息
	Jobs           *[]EncodeServerJob `json:"jobs,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o RestartEncodeServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartEncodeServerResponse struct{}"
	}

	return strings.Join([]string{"RestartEncodeServerResponse", string(data)}, " ")
}
