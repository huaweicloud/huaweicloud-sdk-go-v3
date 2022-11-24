package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddVolumesResponse struct {

	// 增加磁盘任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddVolumesResponse struct{}"
	}

	return strings.Join([]string{"AddVolumesResponse", string(data)}, " ")
}
