package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHostResponse Response Object
type CreateHostResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 主机id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostResponse struct{}"
	}

	return strings.Join([]string{"CreateHostResponse", string(data)}, " ")
}
