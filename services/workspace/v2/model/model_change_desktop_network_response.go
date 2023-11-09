package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDesktopNetworkResponse Response Object
type ChangeDesktopNetworkResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeDesktopNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDesktopNetworkResponse struct{}"
	}

	return strings.Join([]string{"ChangeDesktopNetworkResponse", string(data)}, " ")
}
