package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGroupOrResourceOwnerRequest Request Object
type UpdateGroupOrResourceOwnerRequest struct {
	Body *UpdateResourceOwner `json:"body,omitempty"`
}

func (o UpdateGroupOrResourceOwnerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGroupOrResourceOwnerRequest struct{}"
	}

	return strings.Join([]string{"UpdateGroupOrResourceOwnerRequest", string(data)}, " ")
}
