package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Status 工作项状态
type Status struct {

	// **参数解释：** 状态名称 **取值范围：** 新建 进行中 已解决 测试中 已关闭 已拒绝
	Name *string `json:"name,omitempty"`

	// **参数解释：** 状态id。 **取值范围：** 1（新建） 2（进行中） 3（已解决） 4（测试中） 5（ 已关闭） 6（已拒绝）
	Id *int32 `json:"id,omitempty"`
}

func (o Status) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Status struct{}"
	}

	return strings.Join([]string{"Status", string(data)}, " ")
}
