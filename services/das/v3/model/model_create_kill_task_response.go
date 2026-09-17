package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKillTaskResponse Response Object
type CreateKillTaskResponse struct {

	// 操作是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreateKillTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKillTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateKillTaskResponse", string(data)}, " ")
}
