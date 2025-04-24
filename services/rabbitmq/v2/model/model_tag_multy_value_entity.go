package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagMultyValueEntity struct {

	// **参数解释**： 标签键。 **取值范围**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 标签值。
	Values *[]string `json:"values,omitempty"`
}

func (o TagMultyValueEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagMultyValueEntity struct{}"
	}

	return strings.Join([]string{"TagMultyValueEntity", string(data)}, " ")
}
