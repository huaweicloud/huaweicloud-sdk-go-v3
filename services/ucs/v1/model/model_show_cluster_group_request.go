package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterGroupRequest Request Object
type ShowClusterGroupRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`
}

func (o ShowClusterGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterGroupRequest", string(data)}, " ")
}
