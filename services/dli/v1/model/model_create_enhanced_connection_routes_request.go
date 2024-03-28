package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnhancedConnectionRoutesRequest Request Object
type CreateEnhancedConnectionRoutesRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Body *CreateEnhancedConnectionRoutesRequestBody `json:"body,omitempty"`
}

func (o CreateEnhancedConnectionRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnhancedConnectionRoutesRequest struct{}"
	}

	return strings.Join([]string{"CreateEnhancedConnectionRoutesRequest", string(data)}, " ")
}
