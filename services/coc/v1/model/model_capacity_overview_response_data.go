package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CapacityOverviewResponseData struct {

	// **参数解释：** 硬盘大小总量。 **取值范围：** 云服务对应的总内存。
	SumSize *string `json:"sum_size,omitempty"`

	// **参数解释：** CPU分配量总量。 **取值范围：** 不涉及。
	SumCpu *string `json:"sum_cpu,omitempty"`

	// **参数解释：** 内存分配量总量。 **取值范围：** 不涉及。
	SumMem *string `json:"sum_mem,omitempty"`

	// **参数解释：** 云服务类型。 **取值范围：** 不涉及。
	Provider *string `json:"provider,omitempty"`

	// **参数解释：** 云服务资源类型。。 **取值范围：** 资源类型较多，根据实际业务选择资源类型、常用资源类型如下： - cloudservers：弹性云服务器。 - servers：裸金属服务器。 - clusters：云容器引擎。 - instances：云数据库。
	Type *string `json:"type,omitempty"`
}

func (o CapacityOverviewResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacityOverviewResponseData struct{}"
	}

	return strings.Join([]string{"CapacityOverviewResponseData", string(data)}, " ")
}
