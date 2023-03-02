package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteDbUserResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDbUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteDbUserResponse", string(data)}, " ")
}
