package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSslSwitchResponse Response Object
type UpdateSslSwitchResponse struct {

	// DCS任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateSslSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSslSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateSslSwitchResponse", string(data)}, " ")
}
