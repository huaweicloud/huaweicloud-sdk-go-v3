package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostGroupPermissionsResponse Response Object
type ListHostGroupPermissionsResponse struct {

	// 主机集群权限矩阵
	Body           *[]DevUcClusterPermission `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListHostGroupPermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostGroupPermissionsResponse struct{}"
	}

	return strings.Join([]string{"ListHostGroupPermissionsResponse", string(data)}, " ")
}
