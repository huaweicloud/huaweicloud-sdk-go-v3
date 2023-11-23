package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthorizationPrivilegesRequest Request Object
type ListAuthorizationPrivilegesRequest struct {

	// 授权对象，和授权接口中的object对应 \"jobs.flink.flink作业ID\"，查询指定的作业。 \"groups.程序包组名\"，查询指定的程序包组。 \"resources.程序包名\"，查询指定程序包。
	Object string `json:"object"`
}

func (o ListAuthorizationPrivilegesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizationPrivilegesRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorizationPrivilegesRequest", string(data)}, " ")
}
