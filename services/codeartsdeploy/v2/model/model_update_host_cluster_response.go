package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostClusterResponse Response Object
type UpdateHostClusterResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 主机集群id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateHostClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostClusterResponse struct{}"
	}

	return strings.Join([]string{"UpdateHostClusterResponse", string(data)}, " ")
}
