package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMsdtcLocalHostResponse Response Object
type DeleteMsdtcLocalHostResponse struct {

	// 任务id
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteMsdtcLocalHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMsdtcLocalHostResponse struct{}"
	}

	return strings.Join([]string{"DeleteMsdtcLocalHostResponse", string(data)}, " ")
}
