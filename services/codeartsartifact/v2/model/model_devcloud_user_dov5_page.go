package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DevcloudUserDov5Page struct {

	// **参数解释**: 总记录数。 **取值范围**: 不涉及。
	TotalRecords *int32 `json:"total_records,omitempty"`

	// **参数解释**: 总页数。 **取值范围**: 不涉及。
	TotalPages *int32 `json:"total_pages,omitempty"`

	// **参数解释**: 用户列表。 **取值范围**: 不涉及。
	Data *[]DevcloudUserDov5 `json:"data,omitempty"`
}

func (o DevcloudUserDov5Page) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevcloudUserDov5Page struct{}"
	}

	return strings.Join([]string{"DevcloudUserDov5Page", string(data)}, " ")
}
