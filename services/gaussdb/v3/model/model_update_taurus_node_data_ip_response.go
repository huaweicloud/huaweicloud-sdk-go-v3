package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTaurusNodeDataIpResponse Response Object
type UpdateTaurusNodeDataIpResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTaurusNodeDataIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTaurusNodeDataIpResponse struct{}"
	}

	return strings.Join([]string{"UpdateTaurusNodeDataIpResponse", string(data)}, " ")
}
