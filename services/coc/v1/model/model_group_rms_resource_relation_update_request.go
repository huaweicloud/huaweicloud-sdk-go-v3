package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupRmsResourceRelationUpdateRequest struct {

	// **参数解释：** 资源关联uuid列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及
	IdList []string `json:"id_list"`

	// **参数解释：** 分组id。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及
	GroupId string `json:"group_id"`
}

func (o GroupRmsResourceRelationUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupRmsResourceRelationUpdateRequest struct{}"
	}

	return strings.Join([]string{"GroupRmsResourceRelationUpdateRequest", string(data)}, " ")
}
