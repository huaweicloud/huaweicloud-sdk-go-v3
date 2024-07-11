package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostResponse Response Object
type DeleteHostResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 主机id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostResponse struct{}"
	}

	return strings.Join([]string{"DeleteHostResponse", string(data)}, " ")
}
