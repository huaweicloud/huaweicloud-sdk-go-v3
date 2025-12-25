package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFileDetailByFullNameRequest Request Object
type ShowFileDetailByFullNameRequest struct {

	// **参数解释**： 文件名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FileName string `json:"file_name"`
}

func (o ShowFileDetailByFullNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFileDetailByFullNameRequest struct{}"
	}

	return strings.Join([]string{"ShowFileDetailByFullNameRequest", string(data)}, " ")
}
