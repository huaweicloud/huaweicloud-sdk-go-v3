package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CountPreoccupyIpNumResponse struct {
	PreoccupyIp *PreoccupyIp `json:"preoccupy_ip,omitempty" xml:"preoccupy_ip"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CountPreoccupyIpNumResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPreoccupyIpNumResponse struct{}"
	}

	return strings.Join([]string{"CountPreoccupyIpNumResponse", string(data)}, " ")
}
