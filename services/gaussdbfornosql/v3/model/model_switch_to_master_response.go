package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchToMasterResponse Response Object
type SwitchToMasterResponse struct {

	// 容灾实例主备倒换的工作ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SwitchToMasterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchToMasterResponse struct{}"
	}

	return strings.Join([]string{"SwitchToMasterResponse", string(data)}, " ")
}
