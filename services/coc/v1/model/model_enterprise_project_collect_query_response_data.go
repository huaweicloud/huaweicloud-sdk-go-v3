package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnterpriseProjectCollectQueryResponseData struct {

	// **参数解释：** 唯一标识id。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 用户id。 **取值范围：** 不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释：** 企业项目收藏列表。 **取值范围：** 不涉及。
	EpIdList *[]string `json:"ep_id_list,omitempty"`
}

func (o EnterpriseProjectCollectQueryResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnterpriseProjectCollectQueryResponseData struct{}"
	}

	return strings.Join([]string{"EnterpriseProjectCollectQueryResponseData", string(data)}, " ")
}
