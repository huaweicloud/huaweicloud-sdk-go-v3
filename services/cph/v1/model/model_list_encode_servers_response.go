package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListEncodeServersResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`

	// 编码服务信息
	EncodeServers  []interface{} `json:"encode_servers" xml:"encode_servers"`
	HttpStatusCode int           `json:"-"`
}

func (o ListEncodeServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEncodeServersResponse struct{}"
	}

	return strings.Join([]string{"ListEncodeServersResponse", string(data)}, " ")
}
