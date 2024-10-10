package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReplicationJobResponse Response Object
type DeleteReplicationJobResponse struct {

	// 空响应体。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteReplicationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReplicationJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteReplicationJobResponse", string(data)}, " ")
}
