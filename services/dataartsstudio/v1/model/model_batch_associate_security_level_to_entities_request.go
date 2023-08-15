package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssociateSecurityLevelToEntitiesRequest Request Object
type BatchAssociateSecurityLevelToEntitiesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *BulkSecurityLevel `json:"body,omitempty"`
}

func (o BatchAssociateSecurityLevelToEntitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateSecurityLevelToEntitiesRequest struct{}"
	}

	return strings.Join([]string{"BatchAssociateSecurityLevelToEntitiesRequest", string(data)}, " ")
}
