package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterManagerAuthStateResponse Response Object
type ListClusterManagerAuthStateResponse struct {

	// 查询指定集群界面授权类型，0为界面普通授权，当前仅开放普通授权
	AuthType float32 `json:"auth_type,omitempty"`

	// 查询集群是否已开启界面授权
	Authorized *bool `json:"authorized,omitempty"`

	// 集群界面授权的过期时间
	ExpirationTime *string `json:"expiration_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListClusterManagerAuthStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterManagerAuthStateResponse struct{}"
	}

	return strings.Join([]string{"ListClusterManagerAuthStateResponse", string(data)}, " ")
}
