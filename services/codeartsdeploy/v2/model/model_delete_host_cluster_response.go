package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostClusterResponse Response Object
type DeleteHostClusterResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	// 主机集群id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteHostClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteHostClusterResponse", string(data)}, " ")
}
