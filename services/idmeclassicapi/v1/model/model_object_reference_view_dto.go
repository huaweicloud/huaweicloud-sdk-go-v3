package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectReferenceViewDto struct {

	// **参数解释：**  唯一标识，标识关联实例的主键ID。  **默认取值：**  不涉及。
	Id string `json:"id"`

	// **参数解释：**  类名，标识关联实例的类类型。  **默认取值：**  不涉及。
	Clazz *string `json:"clazz,omitempty"`
}

func (o ObjectReferenceViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectReferenceViewDto struct{}"
	}

	return strings.Join([]string{"ObjectReferenceViewDto", string(data)}, " ")
}
