package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupRmsResourceRelationRequest Request Object
type UpdateGroupRmsResourceRelationRequest struct {
	Body *GroupRmsResourceRelationUpdateRequest `json:"body,omitempty"`
}

func (o UpdateGroupRmsResourceRelationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupRmsResourceRelationRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupRmsResourceRelationRequest", string(data)}, " ")
}
