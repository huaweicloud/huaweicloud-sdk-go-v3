package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// V1DiskExtExpandReq 磁盘扩容检查请求体。
type V1DiskExtExpandReq struct {

	// **参数解释**： 磁盘扩容后单节点存储容量。 **约束限制**： 不涉及。 **取值范围**： 大于等于10。 **默认取值**： 不涉及。
	NewSize *int32 `json:"new_size,omitempty"`
}

func (o V1DiskExtExpandReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V1DiskExtExpandReq struct{}"
	}

	return strings.Join([]string{"V1DiskExtExpandReq", string(data)}, " ")
}
