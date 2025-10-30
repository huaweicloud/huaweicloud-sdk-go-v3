package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulHostHostsResponseInfoVulIdsInfo **参数解释**: 各漏洞类型漏洞id列表
type VulHostHostsResponseInfoVulIdsInfo struct {

	// **参数解释**: Linux漏洞的漏洞id列表
	LinuxVulIdList *[]string `json:"linux_vul_id_list,omitempty"`

	// **参数解释**: Windows漏洞的漏洞id列表
	WindowsVulIdList *[]string `json:"windows_vul_id_list,omitempty"`

	// **参数解释**: Web-CMS漏洞的漏洞id列表
	WebcmsVulIdList *[]string `json:"webcms_vul_id_list,omitempty"`

	// **参数解释**: 应用漏洞的漏洞id列表
	AppVulIdList *[]string `json:"app_vul_id_list,omitempty"`

	// **参数解释**: 应急漏洞的漏洞id列表
	UrgentVulIdList *[]string `json:"urgent_vul_id_list,omitempty"`
}

func (o VulHostHostsResponseInfoVulIdsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulHostHostsResponseInfoVulIdsInfo struct{}"
	}

	return strings.Join([]string{"VulHostHostsResponseInfoVulIdsInfo", string(data)}, " ")
}
