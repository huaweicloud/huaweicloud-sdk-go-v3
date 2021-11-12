package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Org struct {
	// 组织MSP标识

	OrgMspId *string `json:"org_msp_id,omitempty"`
	// 组织域名

	OrgDomain *string `json:"org_domain,omitempty"`
	// key:节点名称，value：节点详细信息

	Peers map[string]Node `json:"peers,omitempty"`
}

func (o Org) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Org struct{}"
	}

	return strings.Join([]string{"Org", string(data)}, " ")
}
