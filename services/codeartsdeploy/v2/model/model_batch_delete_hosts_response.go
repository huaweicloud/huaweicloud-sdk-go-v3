package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteHostsResponse Response Object
type BatchDeleteHostsResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 返回结果
	Result         *[]string `json:"result,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteHostsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteHostsResponse", string(data)}, " ")
}
