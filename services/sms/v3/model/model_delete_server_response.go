package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServerResponse Response Object
type DeleteServerResponse struct {

	// 删除指定ID的源端服务器信息成功
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerResponse", string(data)}, " ")
}
