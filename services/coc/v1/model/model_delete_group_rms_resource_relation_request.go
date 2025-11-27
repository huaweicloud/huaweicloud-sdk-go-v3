package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupRmsResourceRelationRequest Request Object
type DeleteGroupRmsResourceRelationRequest struct {

	// **参数解释：** 分组id列表。 **约束限制：** 不涉及。 **取值范围：** 用户所选的资源所在的分组id组成的list。 **默认取值：** 不涉及。
	IdList []string `json:"id_list"`
}

func (o DeleteGroupRmsResourceRelationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupRmsResourceRelationRequest struct{}"
	}

	return strings.Join([]string{"DeleteGroupRmsResourceRelationRequest", string(data)}, " ")
}
