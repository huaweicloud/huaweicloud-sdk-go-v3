package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGenerationTaskResponse Response Object
type CreateGenerationTaskResponse struct {

	// 分子生成任务ID
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGenerationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGenerationTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateGenerationTaskResponse", string(data)}, " ")
}
