package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSynthesisTaskResponse Response Object
type CreateSynthesisTaskResponse struct {

	// 分子合成路径规划任务ID
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSynthesisTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSynthesisTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateSynthesisTaskResponse", string(data)}, " ")
}
