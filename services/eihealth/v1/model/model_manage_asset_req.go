package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 操作资产请求体
type ManageAssetReq struct {
	Action *AssetAction `json:"action"`
}

func (o ManageAssetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManageAssetReq struct{}"
	}

	return strings.Join([]string{"ManageAssetReq", string(data)}, " ")
}
