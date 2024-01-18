package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeServerImageResponse Response Object
type ChangeServerImageResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeServerImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerImageResponse struct{}"
	}

	return strings.Join([]string{"ChangeServerImageResponse", string(data)}, " ")
}
