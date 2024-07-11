package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHostClusterResponse Response Object
type CreateHostClusterResponse struct {

	// 主机集群id
	Id *string `json:"id,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateHostClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHostClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateHostClusterResponse", string(data)}, " ")
}
