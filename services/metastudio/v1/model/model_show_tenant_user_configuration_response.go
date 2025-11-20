package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantUserConfigurationResponse Response Object
type ShowTenantUserConfigurationResponse struct {

	// 我的收藏列表
	MyCollections *[]CollectionInfo `json:"my_collections,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTenantUserConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantUserConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantUserConfigurationResponse", string(data)}, " ")
}
