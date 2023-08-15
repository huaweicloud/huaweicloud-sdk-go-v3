package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssociateClassificationToEntitiesRequest Request Object
type BatchAssociateClassificationToEntitiesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *OpenBulkClassifications `json:"body,omitempty"`
}

func (o BatchAssociateClassificationToEntitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateClassificationToEntitiesRequest struct{}"
	}

	return strings.Join([]string{"BatchAssociateClassificationToEntitiesRequest", string(data)}, " ")
}
