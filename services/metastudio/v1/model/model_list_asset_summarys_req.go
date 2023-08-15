package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAssetSummarysReq struct {

	// 需要查询的资产ID。
	AssetIds []string `json:"asset_ids"`
}

func (o ListAssetSummarysReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetSummarysReq struct{}"
	}

	return strings.Join([]string{"ListAssetSummarysReq", string(data)}, " ")
}
