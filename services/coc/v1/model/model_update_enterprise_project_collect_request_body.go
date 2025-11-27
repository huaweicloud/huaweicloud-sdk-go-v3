package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEnterpriseProjectCollectRequestBody struct {

	// **参数解释：** 唯一标识id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id string `json:"id"`

	EpIdList []string `json:"ep_id_list"`

	// **参数解释：** 用户id。 **取值范围：** 不涉及。
	UserId *string `json:"user_id,omitempty"`
}

func (o UpdateEnterpriseProjectCollectRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseProjectCollectRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseProjectCollectRequestBody", string(data)}, " ")
}
