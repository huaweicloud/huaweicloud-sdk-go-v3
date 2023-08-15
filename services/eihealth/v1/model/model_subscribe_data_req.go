package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeDataReq 订阅资产市场数据请求体
type SubscribeDataReq struct {

	// 资产ID
	AssetId string `json:"asset_id"`

	// 执行策略（true：全部覆盖，false：全部跳过，默认为true）
	Overwrite *bool `json:"overwrite,omitempty"`

	// 目标文件夹
	TargetFolder *string `json:"target_folder,omitempty"`

	// 版本号
	Version string `json:"version"`
}

func (o SubscribeDataReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeDataReq struct{}"
	}

	return strings.Join([]string{"SubscribeDataReq", string(data)}, " ")
}
