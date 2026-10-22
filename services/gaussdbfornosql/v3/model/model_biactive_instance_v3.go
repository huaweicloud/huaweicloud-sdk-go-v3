package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BiactiveInstanceV3 **参数解释：** 候选实例信息。 **取值范围：** 不涉及。
type BiactiveInstanceV3 struct {

	// **参数解释：** 实例ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 实例名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 实例状态。 **取值范围：** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 实例当前支持的操作。 **取值范围：** 不涉及。
	Actions *[]string `json:"actions,omitempty"`

	// **参数解释：** 项目ID。 **取值范围：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 项目名称。 **取值范围：** 不涉及。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 区域。 **取值范围：** 不涉及。
	RegionCode *string `json:"region_code,omitempty"`

	// **参数解释：** 节点数量。 **取值范围：** 不涉及。
	NodeNum *int32 `json:"node_num,omitempty"`

	// **参数解释：** vCPU数量。 **取值范围：** 不涉及。
	Vcpus *string `json:"vcpus,omitempty"`

	// **参数解释：** 内存大小，单位GB。 **取值范围：** 不涉及。
	Ram *string `json:"ram,omitempty"`

	// **参数解释：** 磁盘大小，单位GB。 **取值范围：** 不涉及。
	DiskSize *string `json:"disk_size,omitempty"`

	// **参数解释：** 已使用磁盘大小，单位GB。 **取值范围：** 不涉及。
	UsedDiskSize *string `json:"used_disk_size,omitempty"`

	// **参数解释：** VPC CIDR。 **取值范围：** 不涉及。
	Cidr *string `json:"cidr,omitempty"`
}

func (o BiactiveInstanceV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BiactiveInstanceV3 struct{}"
	}

	return strings.Join([]string{"BiactiveInstanceV3", string(data)}, " ")
}
