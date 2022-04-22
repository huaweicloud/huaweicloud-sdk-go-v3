package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TranscodeSpecCount struct {

	// 转码规格，具体格式如下： - 若未开启高清低码，则格式为：编码格式_分辨率档位。 - 若已开启高清低码，则格式为：编码格式_PVC_分辨率档位。  其中，编码格式包括H264、H265，分辨率档位包括： - 4K：3840 x 2160及以下 - 2K：2560 x 1440及以下 - FHD：1920 x 1080及以下 - HD：1280 x 720及以下 - SD：640 x 480及以下  示例：若编码格式为H264，分辨率档位为FHD，则转码规格为H264_FHD。
	Type *string `json:"type,omitempty"`

	// 采样时间点转码任务数。
	Count *int64 `json:"count,omitempty"`
}

func (o TranscodeSpecCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TranscodeSpecCount struct{}"
	}

	return strings.Join([]string{"TranscodeSpecCount", string(data)}, " ")
}
