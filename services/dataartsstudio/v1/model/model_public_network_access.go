package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublicNetworkAccess 公网访问诊断结果
type PublicNetworkAccess struct {
	Result *DiagnoseResult `json:"result,omitempty"`

	// kerberos信息
	PublicNetworkInfo *[]PublicNetworkStatus `json:"public_network_info,omitempty"`
}

func (o PublicNetworkAccess) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicNetworkAccess struct{}"
	}

	return strings.Join([]string{"PublicNetworkAccess", string(data)}, " ")
}
