package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConnectionRequest Request Object
type DeleteConnectionRequest struct {

	// 连接ID。
	ConnectionId string `json:"connection_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o DeleteConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteConnectionRequest", string(data)}, " ")
}
