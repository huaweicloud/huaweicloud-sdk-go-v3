package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolTag 资源类型的标签
type PoolTag struct {

	// **参数解释**： 资源类型的标签。 **取值范围**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 资源类型的标签值。
	Values *[]string `json:"values,omitempty"`
}

func (o PoolTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolTag struct{}"
	}

	return strings.Join([]string{"PoolTag", string(data)}, " ")
}
