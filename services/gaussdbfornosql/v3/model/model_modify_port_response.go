package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ModifyPortResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyPortResponse struct{}"
	}

	return strings.Join([]string{"ModifyPortResponse", string(data)}, " ")
}
