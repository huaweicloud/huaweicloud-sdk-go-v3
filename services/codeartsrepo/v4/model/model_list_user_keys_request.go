package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserKeysRequest Request Object
type ListUserKeysRequest struct {

	// **参数解释：** 偏移量，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 返回数量。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** key的标题名称。 **取值范围：** 字符串长度不少于1，不超过2000。
	Query *string `json:"query,omitempty"`
}

func (o ListUserKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserKeysRequest struct{}"
	}

	return strings.Join([]string{"ListUserKeysRequest", string(data)}, " ")
}
