package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagResp **参数解释**： 标签列表。
type TagResp struct {

	// **参数解释**： 标签名。 **取值范围**： 最大长度128个unicode字符。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 标签值。 **取值范围**： 每个值最大长度255个unicode字符。如果values为空列表，则表示查询任意value
	Values *[]string `json:"values,omitempty"`
}

func (o TagResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResp struct{}"
	}

	return strings.Join([]string{"TagResp", string(data)}, " ")
}
