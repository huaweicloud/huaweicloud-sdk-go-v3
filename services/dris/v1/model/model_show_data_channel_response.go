package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDataChannelResponse struct {

	// **参数说明**：平台类型。   **取值范围**：  - DRIS：华为路网数字化平台  - LITONG：利通  - ZHONGQIYAN：中汽研
	PlatformType *string `json:"platform_type,omitempty"`

	PlatformPara *PlatformPara `json:"platform_para,omitempty"`

	// **参数说明**：华为路网数字化平台或第三方业务平台连接状态。  **取值范围**： - ONLINE：在线 - OFFLINE：离线 - INITIAL：初始化
	ChannelStatus  *string `json:"channel_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDataChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataChannelResponse struct{}"
	}

	return strings.Join([]string{"ShowDataChannelResponse", string(data)}, " ")
}
