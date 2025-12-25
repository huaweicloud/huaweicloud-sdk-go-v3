package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataClassRef struct {

	// **参数解释**: 数据类的ID **取值范围**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 数据类的名称 **取值范围**: 不涉及
	Name string `json:"name"`
}

func (o DataClassRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassRef struct{}"
	}

	return strings.Join([]string{"DataClassRef", string(data)}, " ")
}
