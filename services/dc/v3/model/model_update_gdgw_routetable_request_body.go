package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateGdgwRoutetableRequestBody struct {

	// 空运行：true-是，false-否
	DryRun *bool `json:"dry_run,omitempty"`

	GdgwRoutetable *UpdateGdgwRoutetableRequestBodyGdgwRoutetable `json:"gdgw_routetable,omitempty"`
}

func (o UpdateGdgwRoutetableRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGdgwRoutetableRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGdgwRoutetableRequestBody", string(data)}, " ")
}
