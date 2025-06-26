package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServersResponse Response Object
type DeleteServersResponse struct {

	// 批量删除源端服务器的返回
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServersResponse struct{}"
	}

	return strings.Join([]string{"DeleteServersResponse", string(data)}, " ")
}
