package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BasicResourceTagResponseData struct {

	// **参数解释：** 标签id。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`
}

func (o BasicResourceTagResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BasicResourceTagResponseData struct{}"
	}

	return strings.Join([]string{"BasicResourceTagResponseData", string(data)}, " ")
}
