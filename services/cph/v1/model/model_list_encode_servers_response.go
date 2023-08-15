package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEncodeServersResponse Response Object
type ListEncodeServersResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 编码服务信息。
	EncodeServers  *[]EncodeServer `json:"encode_servers,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListEncodeServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEncodeServersResponse struct{}"
	}

	return strings.Join([]string{"ListEncodeServersResponse", string(data)}, " ")
}
