package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyConnectionPoolTypeResponse Response Object
type UpdateProxyConnectionPoolTypeResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateProxyConnectionPoolTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyConnectionPoolTypeResponse struct{}"
	}

	return strings.Join([]string{"UpdateProxyConnectionPoolTypeResponse", string(data)}, " ")
}
