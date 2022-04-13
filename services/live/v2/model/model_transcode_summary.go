package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TranscodeSummary struct {
	// 转码规格，格式是“编码格式_分辨率档位”（未开启高清低码）和“编码格式_PVC_分辨率档位”（开启高清低码）。  其中编码格式包括H264、H265，分辨率档位包括：  4K（3840 x 2160）及以下，2K（2560 x 1440）及以下，FHD（1920 x 1080）及以下，HD（1280 x 720）及以下，SD（640 x 480）及以下。

	Type *string `json:"type,omitempty"`
	// 总转码时长，单位为分钟，保留两位小数。

	Value *float64 `json:"value,omitempty"`
}

func (o TranscodeSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscodeSummary struct{}"
	}

	return strings.Join([]string{"TranscodeSummary", string(data)}, " ")
}
