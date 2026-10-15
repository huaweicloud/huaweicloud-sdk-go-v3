package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceSpecResponse **参数解释**：资源规格配置，用于描述Driver或Executor的CPU、内存、磁盘资源配置。
type ShowResourceSpecResponse struct {

	// **参数解释**：CPU毫核数，用于指定Driver或Executor的CPU资源。 **取值范围**：最小值为1,默认值为1000。
	Cpu *int32 `json:"cpu,omitempty"`

	// **参数解释**：内存大小，用于指定Driver或Executor的内存资源。 **取值范围**：单位为MB，最小值为1。默认值为4096。
	Memory *int32 `json:"memory,omitempty"`

	// **参数解释**：本地磁盘大小，用于指定Driver或Executor的本地磁盘容量。 **取值范围**：单位为GB，最小值为0。
	Disk *int32 `json:"disk,omitempty"`
}

func (o ShowResourceSpecResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceSpecResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceSpecResponse", string(data)}, " ")
}
