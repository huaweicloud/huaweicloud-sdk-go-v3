package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLdapConfigResponse Response Object
type CreateLdapConfigResponse struct {

	// ldap异步任务的id。可通过查询job的状态详情接口查询job的执行状态。
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLdapConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLdapConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateLdapConfigResponse", string(data)}, " ")
}
