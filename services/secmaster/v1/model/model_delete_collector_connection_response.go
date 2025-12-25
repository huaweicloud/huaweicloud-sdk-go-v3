package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCollectorConnectionResponse Response Object
type DeleteCollectorConnectionResponse struct {

	// 创建成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteCollectorConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCollectorConnectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteCollectorConnectionResponse", string(data)}, " ")
}
