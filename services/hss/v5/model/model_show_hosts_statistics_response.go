package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostsStatisticsResponse Response Object
type ShowHostsStatisticsResponse struct {

	// 服务器总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 有风险服务器数
	RiskNum *int32 `json:"risk_num,omitempty"`

	// 未防护服务器数
	UnprotectedNum *int32 `json:"unprotected_num,omitempty"`

	// Agent未安装数
	NotInstalledNum *int32 `json:"not_installed_num,omitempty"`

	// Agent安装失败数
	InstalledFailedNum *int32 `json:"installed_failed_num,omitempty"`

	// Agent不在线数
	NotOnlineNum *int32 `json:"not_online_num,omitempty"`

	// 开启基础版服务器数
	VersionBasicNum *int32 `json:"version_basic_num,omitempty"`

	// 开启专业版服务器数
	VersionAdvancedNum *int32 `json:"version_advanced_num,omitempty"`

	// 开启企业版服务器数
	VersionEnterpriseNum *int32 `json:"version_enterprise_num,omitempty"`

	// 开启旗舰版服务器数
	VersionPremiumNum *int32 `json:"version_premium_num,omitempty"`

	// 开启网页防篡改版服务器数
	VersionWtpNum *int32 `json:"version_wtp_num,omitempty"`

	// 开启容器版服务器数
	VersionContainerNum *int32 `json:"version_container_num,omitempty"`

	// 服务器组总数
	HostGroupNum *int32 `json:"host_group_num,omitempty"`

	// 资产服务器组数量
	ServerGroupNum *int32 `json:"server_group_num,omitempty"`

	// 资产重要性列表
	AssetValueList *[]AssetValueHostNumInfo `json:"asset_value_list,omitempty"`

	// 服务器组列表
	ServerGroupList *[]ServerGroupItem `json:"server_group_list,omitempty"`

	// 已忽略服务器数
	IgnoreHostNum *int32 `json:"ignore_host_num,omitempty"`

	// 防护中服务器数
	ProtectedNum *int32 `json:"protected_num,omitempty"`

	// 防护中断服务器数
	ProtectInterruptNum *int32 `json:"protect_interrupt_num,omitempty"`

	// 空闲配额数
	IdleNum *int32 `json:"idle_num,omitempty"`

	// 旗舰版主机未开启agent自保护数
	PremiumNonSpNum *int32 `json:"premium_non_sp_num,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ShowHostsStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostsStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowHostsStatisticsResponse", string(data)}, " ")
}
