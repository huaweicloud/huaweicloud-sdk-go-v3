package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateProtectedInstanceResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProtectedInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProtectedInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateProtectedInstanceResponse", string(data)}, " ")
}
