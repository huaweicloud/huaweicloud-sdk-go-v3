package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeImageReq 订阅镜像实体
type SubscribeImageReq struct {

	// 镜像资产ID
	AssetId string `json:"asset_id"`

	// 镜像资产版本
	Version string `json:"version"`
}

func (o SubscribeImageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeImageReq struct{}"
	}

	return strings.Join([]string{"SubscribeImageReq", string(data)}, " ")
}
