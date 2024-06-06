package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartProxyInstanceResponse Response Object
type RestartProxyInstanceResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartProxyInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartProxyInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestartProxyInstanceResponse", string(data)}, " ")
}
