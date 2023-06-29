package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSearchTaskResponse Response Object
type CreateSearchTaskResponse struct {

	// 分子搜索任务ID
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSearchTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSearchTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateSearchTaskResponse", string(data)}, " ")
}
