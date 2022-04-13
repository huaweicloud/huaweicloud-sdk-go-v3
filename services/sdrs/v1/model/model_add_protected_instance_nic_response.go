package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddProtectedInstanceNicResponse struct {
	// 成功返回jobId信息

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddProtectedInstanceNicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectedInstanceNicResponse struct{}"
	}

	return strings.Join([]string{"AddProtectedInstanceNicResponse", string(data)}, " ")
}
