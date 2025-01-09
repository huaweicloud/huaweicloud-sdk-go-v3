package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerResponse Response Object
type DeleteServerResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerResponse", string(data)}, " ")
}
