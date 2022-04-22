package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Environment struct {

	// 环境供应商，HWCP/HWC/AWS/Azure/GCP等。
	Type *string `json:"type,omitempty"`

	// 租户账号ID，用来标识事件所属租户。
	DomainId *string `json:"domain_id,omitempty"`

	// 租户项目ID，用来标识事件所属项目区域。
	ProjectId *string `json:"project_id,omitempty"`

	// 数据源产品所在区域，具体取值范围查看华为云地区和终端节点定义。
	RegionId *string `json:"region_id,omitempty"`
}

func (o Environment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Environment struct{}"
	}

	return strings.Join([]string{"Environment", string(data)}, " ")
}
