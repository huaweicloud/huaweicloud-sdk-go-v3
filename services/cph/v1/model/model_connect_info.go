package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectInfo 云手机信息。
type ConnectInfo struct {

	// 云手机的唯一标识。
	PhoneId *string `json:"phone_id,omitempty"`

	AccessInfo *ConnectInfoAccessInfo `json:"access_info,omitempty"`
}

func (o ConnectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectInfo struct{}"
	}

	return strings.Join([]string{"ConnectInfo", string(data)}, " ")
}
