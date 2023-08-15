package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGaussMySqlProxyListResponse Response Object
type ShowGaussMySqlProxyListResponse struct {

	// Proxy实例信息列表。
	ProxyList      *[]MysqlShowProxyResponseV3 `json:"proxy_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowGaussMySqlProxyListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlProxyListResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlProxyListResponse", string(data)}, " ")
}
