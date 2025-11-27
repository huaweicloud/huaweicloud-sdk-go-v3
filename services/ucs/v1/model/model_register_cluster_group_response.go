package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterClusterGroupResponse Response Object
type RegisterClusterGroupResponse struct {

	// 容器舰队的UID信息
	Uid            *string `json:"uid,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RegisterClusterGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterClusterGroupResponse struct{}"
	}

	return strings.Join([]string{"RegisterClusterGroupResponse", string(data)}, " ")
}
