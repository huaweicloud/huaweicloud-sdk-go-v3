package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProtectedInstanceNicResponse Response Object
type DeleteProtectedInstanceNicResponse struct {

	// 成功返回jobId信息
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProtectedInstanceNicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceNicResponse struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceNicResponse", string(data)}, " ")
}
