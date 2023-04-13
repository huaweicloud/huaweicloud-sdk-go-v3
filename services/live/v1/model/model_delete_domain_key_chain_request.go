package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDomainKeyChainRequest struct {

	// 服务鉴权Token，服务开启鉴权，必须携带Access-Control-Allow-Internal访问服务。
	AccessControlAllowInternal *string `json:"Access-Control-Allow-Internal,omitempty"`

	// 服务鉴权Token，服务开启鉴权，必须携带Access-Control-Allow-External访问服务。
	AccessControlAllowExternal *string `json:"Access-Control-Allow-External,omitempty"`

	// 直播域名，包括推流域名和播放域名
	Domain string `json:"domain"`
}

func (o DeleteDomainKeyChainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainKeyChainRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainKeyChainRequest", string(data)}, " ")
}
