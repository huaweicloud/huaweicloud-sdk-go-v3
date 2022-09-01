package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 漏洞列表
type VulInfo struct {

	// 漏洞名称
	VulName *string `json:"vul_name,omitempty" xml:"vul_name"`

	// 漏洞ID
	VulId *string `json:"vul_id,omitempty" xml:"vul_id"`

	// 漏洞标签
	LabelList *[]string `json:"label_list,omitempty" xml:"label_list"`

	// 修复必要性
	RepairNecessity *int32 `json:"repair_necessity,omitempty" xml:"repair_necessity"`

	// 受影响服务器台数
	HostNum *int32 `json:"host_num,omitempty" xml:"host_num"`

	// 未处理服务器台数
	UnhandleHostNum *int32 `json:"unhandle_host_num,omitempty" xml:"unhandle_host_num"`

	// 最近扫描时间
	ScanTime *int64 `json:"scan_time,omitempty" xml:"scan_time"`

	// 解决方案
	SolutionDetail *string `json:"solution_detail,omitempty" xml:"solution_detail"`

	// URL链接
	Url *string `json:"url,omitempty" xml:"url"`

	// 漏洞描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 漏洞类型，包含如下：   -linux_vul : linux漏洞   -windows_vul : windows漏洞   -web_cms : Web-CMS漏洞
	Type *string `json:"type,omitempty" xml:"type"`

	// 主机列表
	HostIdList *[]string `json:"host_id_list,omitempty" xml:"host_id_list"`
}

func (o VulInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulInfo struct{}"
	}

	return strings.Join([]string{"VulInfo", string(data)}, " ")
}
