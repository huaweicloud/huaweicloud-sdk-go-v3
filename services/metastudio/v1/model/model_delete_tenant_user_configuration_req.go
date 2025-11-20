package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTenantUserConfigurationReq 删除租户个性化配置请求
type DeleteTenantUserConfigurationReq struct {

	// 我的收藏列表
	MyCollections *[]CollectionInfo `json:"my_collections,omitempty"`
}

func (o DeleteTenantUserConfigurationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantUserConfigurationReq struct{}"
	}

	return strings.Join([]string{"DeleteTenantUserConfigurationReq", string(data)}, " ")
}
