package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProjectCreatorDto struct {

	// **参数解释：** 唯一标识id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 名称。 **取值范围：** 字符串长度不少于1，不超过1000。
	Name *string `json:"name,omitempty"`

	// **参数解释：** iam_id。 **取值范围：** 字符串长度不少于1，不超过1000。
	Username *string `json:"username,omitempty"`
}

func (o ProjectCreatorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProjectCreatorDto struct{}"
	}

	return strings.Join([]string{"ProjectCreatorDto", string(data)}, " ")
}
