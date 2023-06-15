package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateProxySessionConsistenceRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	// 租户在某一instance下的数据库代理ID。
	ProxyId string `json:"proxy_id"`

	Body *ModifyProxyConsistRequest `json:"body,omitempty"`
}

func (o UpdateProxySessionConsistenceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxySessionConsistenceRequest struct{}"
	}

	return strings.Join([]string{"UpdateProxySessionConsistenceRequest", string(data)}, " ")
}
