package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeDesktopNetworkResponse Response Object
type BatchChangeDesktopNetworkResponse struct {

	// 任务id。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchChangeDesktopNetworkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeDesktopNetworkResponse struct{}"
	}

	return strings.Join([]string{"BatchChangeDesktopNetworkResponse", string(data)}, " ")
}
