package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFileDetailRequest Request Object
type ShowFileDetailRequest struct {

	// **参数解释**： 文件id。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id string `json:"id"`
}

func (o ShowFileDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowFileDetailRequest", string(data)}, " ")
}
