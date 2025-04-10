package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulhandleHistoryResponseInfoDataList struct {

	// 历史记录唯一id
	Id *string `json:"id,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 漏洞类型
	Type *string `json:"type,omitempty"`

	// 服务器ID
	HostId *string `json:"host_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器公网ip
	PublicIp *string `json:"public_ip,omitempty"`

	// 服务器私网ip
	PrivateIp *string `json:"private_ip,omitempty"`

	// 处置时间
	HandleTime *string `json:"handle_time,omitempty"`

	// 处置状态
	Status *string `json:"status,omitempty"`

	// 失败原因
	FailedReason *string `json:"failed_reason,omitempty"`

	// 备注
	Description *string `json:"description,omitempty"`

	// 漏洞ID
	VulId *string `json:"vul_id,omitempty"`

	// 漏洞名称
	VulName *string `json:"vul_name,omitempty"`

	// 资产重要性
	AssetValue *string `json:"asset_value,omitempty"`

	// cve列表
	CveList *[]VulCveInfo `json:"cve_list,omitempty"`

	// 软件名称
	AppName *string `json:"app_name,omitempty"`

	// 应用软件路径
	AppPath *string `json:"app_path,omitempty"`

	// 软件版本
	AppVersion *string `json:"app_version,omitempty"`

	// 处置类型
	HandleType *string `json:"handle_type,omitempty"`

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o VulhandleHistoryResponseInfoDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulhandleHistoryResponseInfoDataList struct{}"
	}

	return strings.Join([]string{"VulhandleHistoryResponseInfoDataList", string(data)}, " ")
}
