package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGdgwRoutetableRequestBody 修改全域接入网关路由表的request
type UpdateGdgwRoutetableRequestBody struct {

	// 是否dry run模式执行
	DryRun *bool `json:"dry_run,omitempty"`

	GdgwRoutetable *GdgwRouteTableRequest `json:"gdgw_routetable,omitempty"`
}

func (o UpdateGdgwRoutetableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGdgwRoutetableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGdgwRoutetableRequestBody", string(data)}, " ")
}
