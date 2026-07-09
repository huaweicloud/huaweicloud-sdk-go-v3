package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceEnvironment struct {

	// 账户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 企业项目ID
	EpId *string `json:"ep_id,omitempty"`

	// 企业项目名称
	EpName *string `json:"ep_name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// regionID
	RegionId *string `json:"region_id,omitempty"`
}

func (o ResourceEnvironment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceEnvironment struct{}"
	}

	return strings.Join([]string{"ResourceEnvironment", string(data)}, " ")
}
