package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTenantUserConfigurationReq 设置租户个性化配置请求
type SetTenantUserConfigurationReq struct {

	// 我的收藏列表
	MyCollections *[]CollectionInfo `json:"my_collections,omitempty"`
}

func (o SetTenantUserConfigurationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTenantUserConfigurationReq struct{}"
	}

	return strings.Join([]string{"SetTenantUserConfigurationReq", string(data)}, " ")
}
