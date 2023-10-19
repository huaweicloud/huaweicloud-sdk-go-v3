package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyNewConfigurationsResponse Response Object
type UpdateProxyNewConfigurationsResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateProxyNewConfigurationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyNewConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"UpdateProxyNewConfigurationsResponse", string(data)}, " ")
}
