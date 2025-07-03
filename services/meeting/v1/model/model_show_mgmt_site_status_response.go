package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMgmtSiteStatusResponse Response Object
type ShowMgmtSiteStatusResponse struct {

	// 管理站点状态 NORMAL:正常  FAULT:故障
	Status *string `json:"status,omitempty"`

	// 链接入会跳转url
	RedirectJoinUrl *string `json:"redirectJoinUrl,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowMgmtSiteStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMgmtSiteStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowMgmtSiteStatusResponse", string(data)}, " ")
}
