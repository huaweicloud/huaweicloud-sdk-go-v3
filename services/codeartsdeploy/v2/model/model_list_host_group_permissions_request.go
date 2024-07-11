package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostGroupPermissionsRequest Request Object
type ListHostGroupPermissionsRequest struct {

	// 主机集群id
	GroupId string `json:"group_id"`
}

func (o ListHostGroupPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListHostGroupPermissionsRequest", string(data)}, " ")
}
