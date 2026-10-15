package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFlinkResourceSpecResponseBody **参数解释**：资源规格，包含单个组件（JobManager或TaskManager）的CPU、内存和磁盘配额。
type ShowFlinkResourceSpecResponseBody struct {

	// **参数解释**：CPU毫核数，单位为毫核。 **取值范围**：500~48000。
	Cpu int32 `json:"cpu"`

	// **参数解释**：内存大小，单位为MB。 **取值范围**：2048~393216。
	Memory int32 `json:"memory"`

	// **参数解释**：Pod临时存储容量，单位为GB。 **取值范围**：30~1024。
	Disk int32 `json:"disk"`
}

func (o ShowFlinkResourceSpecResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlinkResourceSpecResponseBody struct{}"
	}

	return strings.Join([]string{"ShowFlinkResourceSpecResponseBody", string(data)}, " ")
}
