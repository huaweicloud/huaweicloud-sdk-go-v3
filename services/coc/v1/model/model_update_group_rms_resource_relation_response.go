package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupRmsResourceRelationResponse Response Object
type UpdateGroupRmsResourceRelationResponse struct {

	// **参数解释：** CloudCMDB分配的分组资源id。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGroupRmsResourceRelationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupRmsResourceRelationResponse struct{}"
	}

	return strings.Join([]string{"UpdateGroupRmsResourceRelationResponse", string(data)}, " ")
}
