package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportHandledVulnerabilitiesRequestBodySpecificVuls struct {

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 漏洞id
	VulId *string `json:"vul_id,omitempty"`
}

func (o ExportHandledVulnerabilitiesRequestBodySpecificVuls) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportHandledVulnerabilitiesRequestBodySpecificVuls struct{}"
	}

	return strings.Join([]string{"ExportHandledVulnerabilitiesRequestBodySpecificVuls", string(data)}, " ")
}
