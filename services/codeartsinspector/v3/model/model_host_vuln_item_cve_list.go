package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostVulnItemCveList struct {

	// CVE漏洞ID
	Id *string `json:"id,omitempty"`

	// CVE漏洞链接
	Link *string `json:"link,omitempty"`
}

func (o HostVulnItemCveList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulnItemCveList struct{}"
	}

	return strings.Join([]string{"HostVulnItemCveList", string(data)}, " ")
}
