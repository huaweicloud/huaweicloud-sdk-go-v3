package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowGaussMySqlProxyListResponse struct {

	// Proxy实例信息列表。
	ProxyList      *[]MysqlShowProxyResponseV3 `json:"proxy_list,omitempty" xml:"proxy_list"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowGaussMySqlProxyListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlProxyListResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlProxyListResponse", string(data)}, " ")
}
