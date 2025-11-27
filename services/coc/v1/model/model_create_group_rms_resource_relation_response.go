package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupRmsResourceRelationResponse Response Object
type CreateGroupRmsResourceRelationResponse struct {

	// **参数解释：** CloudCMDB分配的uuid。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGroupRmsResourceRelationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupRmsResourceRelationResponse struct{}"
	}

	return strings.Join([]string{"CreateGroupRmsResourceRelationResponse", string(data)}, " ")
}
