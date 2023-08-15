package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProxyPortRequest Request Object
type UpdateProxyPortRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 租户在某一project下的实例ID。
	InstanceId string `json:"instance_id"`

	// 租户在某一实下的数据库代理ID。
	ProxyId string `json:"proxy_id"`

	Body *UpdateProxyPortRequestBody `json:"body,omitempty"`
}

func (o UpdateProxyPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProxyPortRequest struct{}"
	}

	return strings.Join([]string{"UpdateProxyPortRequest", string(data)}, " ")
}
