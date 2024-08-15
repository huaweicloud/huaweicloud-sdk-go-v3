package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CdmClusterEnterpriseProject struct {

	// 集群ID
	ClusterId *string `json:"cluster_id,omitempty"`

	// 企业项目列表
	SysTags *[]SysTags `json:"sys_tags,omitempty"`
}

func (o CdmClusterEnterpriseProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CdmClusterEnterpriseProject struct{}"
	}

	return strings.Join([]string{"CdmClusterEnterpriseProject", string(data)}, " ")
}
