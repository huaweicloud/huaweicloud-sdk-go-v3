package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLdapConfigResponse Response Object
type DeleteLdapConfigResponse struct {

	// 创建ldap异步任务的id
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteLdapConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLdapConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteLdapConfigResponse", string(data)}, " ")
}
