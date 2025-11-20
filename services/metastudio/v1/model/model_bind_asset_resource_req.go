package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindAssetResourceReq 绑定资源请求
type BindAssetResourceReq struct {

	// 资产id。
	AssetId *string `json:"asset_id,omitempty"`

	// 业务类型。
	BusinessType *string `json:"business_type,omitempty"`
}

func (o BindAssetResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindAssetResourceReq struct{}"
	}

	return strings.Join([]string{"BindAssetResourceReq", string(data)}, " ")
}
