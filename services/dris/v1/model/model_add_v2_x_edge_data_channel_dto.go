package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddV2XEdgeDataChannelDto Data Channel配置数据
type AddV2XEdgeDataChannelDto struct {

	// **参数说明**：平台类型。  **取值范围**： - DRIS：华为路网数字化平台 - LITONG：利通 - ZHONGQIYAN：中汽研
	PlatformType string `json:"platform_type"`

	PlatformPara *PlatformPara `json:"platform_para,omitempty"`
}

func (o AddV2XEdgeDataChannelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddV2XEdgeDataChannelDto struct{}"
	}

	return strings.Join([]string{"AddV2XEdgeDataChannelDto", string(data)}, " ")
}
