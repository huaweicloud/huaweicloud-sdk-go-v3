package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteProtectedInstanceResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProtectedInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceResponse", string(data)}, " ")
}
