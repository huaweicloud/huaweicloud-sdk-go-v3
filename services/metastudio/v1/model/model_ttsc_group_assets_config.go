package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TtscGroupAssetsConfig struct {

	// 分组id
	GroupId *string `json:"group_id,omitempty"`

	// 分组名称
	GroupName *string `json:"group_name,omitempty"`

	// 每个分组的资产id列表
	AssetIds *[]string `json:"asset_ids,omitempty"`
}

func (o TtscGroupAssetsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtscGroupAssetsConfig struct{}"
	}

	return strings.Join([]string{"TtscGroupAssetsConfig", string(data)}, " ")
}
