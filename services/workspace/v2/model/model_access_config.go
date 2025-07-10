package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessConfig 接入信息。
type AccessConfig struct {

	// 接入方式。 - INTERNET：表示互联网接入 - DEDICATED：表示专线接入 - BOTH：表示同时支持互联网接入和专线接入
	AccessMode *string `json:"access_mode,omitempty"`

	// 互联网接入地址，只有access_mode为“INTERNET”或“BOTH”时才会返回该参数。
	InternetAccessAddress *string `json:"internet_access_address,omitempty"`

	// 互联网接入端口。
	InternetAccessPort *string `json:"internet_access_port,omitempty"`

	// 专线接入地址，只有access_mode为“DEDICATED”或“BOTH”时才会返回该参数。
	DedicatedAccessAddress *string `json:"dedicated_access_address,omitempty"`

	// 专线接入备用地址，只有当开启专线备用线路时才会返回该参数。
	DedicatedAccessStandbyAddress *[]string `json:"dedicated_access_standby_address,omitempty"`

	// 专线备用线路失败错误码。
	StandbyAddressResultCode *string `json:"standby_address_result_code,omitempty"`

	// 专线接入网段。接入模式包含专线方式时选填，多个网段信息用分号隔开，列表长度不超过5。
	DedicatedCidrs *string `json:"dedicated_cidrs,omitempty"`
}

func (o AccessConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfig struct{}"
	}

	return strings.Join([]string{"AccessConfig", string(data)}, " ")
}
