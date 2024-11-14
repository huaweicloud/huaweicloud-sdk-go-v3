package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLdapConfigResponse Response Object
type UpdateLdapConfigResponse struct {

	// ldap异步任务的id。可通过查询job的状态详情接口查询job的执行状态。
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLdapConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLdapConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateLdapConfigResponse", string(data)}, " ")
}
