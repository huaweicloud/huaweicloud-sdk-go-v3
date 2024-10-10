package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReplicationJobResponse Response Object
type UpdateReplicationJobResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateReplicationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReplicationJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateReplicationJobResponse", string(data)}, " ")
}
