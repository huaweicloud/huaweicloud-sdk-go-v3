package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeconfigMeta struct {

	// **参数解释**： 固定为node-config。 **取值范围**： 不涉及。
	Name string `json:"name"`
}

func (o NodeconfigMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeconfigMeta struct{}"
	}

	return strings.Join([]string{"NodeconfigMeta", string(data)}, " ")
}
