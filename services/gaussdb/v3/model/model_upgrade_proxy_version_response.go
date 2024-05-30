package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeProxyVersionResponse Response Object
type UpgradeProxyVersionResponse struct {

	// 数据库代理升级信息列表。
	UpdateResult   *[]ProxyUpgradeVersionDetail `json:"update_result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o UpgradeProxyVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeProxyVersionResponse struct{}"
	}

	return strings.Join([]string{"UpgradeProxyVersionResponse", string(data)}, " ")
}
