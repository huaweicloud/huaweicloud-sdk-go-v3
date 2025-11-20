package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Restore2ExistResponse Response Object
type Restore2ExistResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o Restore2ExistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Restore2ExistResponse struct{}"
	}

	return strings.Join([]string{"Restore2ExistResponse", string(data)}, " ")
}
