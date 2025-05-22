package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGeminiDbDualActiveResponse Response Object
type CreateGeminiDbDualActiveResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGeminiDbDualActiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGeminiDbDualActiveResponse struct{}"
	}

	return strings.Join([]string{"CreateGeminiDbDualActiveResponse", string(data)}, " ")
}
