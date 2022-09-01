package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateClusterResponse struct {
	Cluster *CreateClusterResp `json:"cluster,omitempty" xml:"cluster"`

	// 订单号。只有包周期集群拥有该参数。
	OrdeId         *string `json:"ordeId,omitempty" xml:"ordeId"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterResponse struct{}"
	}

	return strings.Join([]string{"CreateClusterResponse", string(data)}, " ")
}
