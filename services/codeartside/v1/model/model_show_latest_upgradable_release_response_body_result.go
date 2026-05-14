package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLatestUpgradableReleaseResponseBodyResult 返回值
type ShowLatestUpgradableReleaseResponseBodyResult struct {

	// 哈希值
	Hash string `json:"hash"`

	// 版本是否可用
	Available bool `json:"available"`

	// 子产品名称
	SubProductName string `json:"sub_product_name"`

	// 子产品版本
	Version string `json:"version"`

	// 下载链接
	Url string `json:"url"`
}

func (o ShowLatestUpgradableReleaseResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLatestUpgradableReleaseResponseBodyResult struct{}"
	}

	return strings.Join([]string{"ShowLatestUpgradableReleaseResponseBodyResult", string(data)}, " ")
}
