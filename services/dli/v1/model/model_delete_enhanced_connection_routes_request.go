package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnhancedConnectionRoutesRequest Request Object
type DeleteEnhancedConnectionRoutesRequest struct {

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId string `json:"connection_id"`

	Name string `json:"name"`
}

func (o DeleteEnhancedConnectionRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnhancedConnectionRoutesRequest struct{}"
	}

	return strings.Join([]string{"DeleteEnhancedConnectionRoutesRequest", string(data)}, " ")
}
