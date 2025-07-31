package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetTtscGroupAssetsRequestBody struct {

	// 资产id
	AssetIds *[]string `json:"asset_ids,omitempty"`
}

func (o SetTtscGroupAssetsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTtscGroupAssetsRequestBody struct{}"
	}

	return strings.Join([]string{"SetTtscGroupAssetsRequestBody", string(data)}, " ")
}
