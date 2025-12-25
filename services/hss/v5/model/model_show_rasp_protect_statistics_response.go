package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRaspProtectStatisticsResponse Response Object
type ShowRaspProtectStatisticsResponse struct {

	// **参数解释** 当前项目（或指定企业项目）下已添加应用防护的云服务器总数，统计范围为所有已启用RASP防护的主机 **取值范围** 取值0-为该项目下云服务器总数量（无上限，实际受账号资源配额限制）
	ProtectHostNum *int64 `json:"protect_host_num,omitempty"`

	// **参数解释** 近7天内当前项目（或指定企业项目）下RASP防护成功拦截的篡改类攻击总数，与功能介绍中'近七天微服务RASP攻击数量'对应 **取值范围** 取值0-无上限（实际受攻击频次限制）
	AntiTamperingNum *int32 `json:"anti_tampering_num,omitempty"`
	HttpStatusCode   int    `json:"-"`
}

func (o ShowRaspProtectStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRaspProtectStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowRaspProtectStatisticsResponse", string(data)}, " ")
}
