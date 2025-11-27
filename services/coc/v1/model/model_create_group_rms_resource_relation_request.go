package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGroupRmsResourceRelationRequest Request Object
type CreateGroupRmsResourceRelationRequest struct {
	Body *GroupRmsResourceRelationCreateRequest `json:"body,omitempty"`
}

func (o CreateGroupRmsResourceRelationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGroupRmsResourceRelationRequest struct{}"
	}

	return strings.Join([]string{"CreateGroupRmsResourceRelationRequest", string(data)}, " ")
}
