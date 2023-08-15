package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeypairResponse Response Object
type UpdateKeypairResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务信息
	Jobs           *[]ServerJob `json:"jobs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeypairResponse struct{}"
	}

	return strings.Join([]string{"UpdateKeypairResponse", string(data)}, " ")
}
