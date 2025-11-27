package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterGroupRequest Request Object
type DeleteClusterGroupRequest struct {

	// 容器舰队ID
	Clustergroupid string `json:"clustergroupid"`
}

func (o DeleteClusterGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterGroupRequest", string(data)}, " ")
}
