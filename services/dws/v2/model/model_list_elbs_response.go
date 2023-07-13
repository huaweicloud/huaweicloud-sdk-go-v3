package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListElbsResponse Response Object
type ListElbsResponse struct {

	// 弹性负载均衡列表
	Elbs           *[]ClusterElbInfo `json:"elbs,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListElbsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListElbsResponse struct{}"
	}

	return strings.Join([]string{"ListElbsResponse", string(data)}, " ")
}
