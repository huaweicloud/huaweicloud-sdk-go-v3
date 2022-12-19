package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDataobjectRelationRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// type of dataclass
	DataclassType string `json:"dataclass_type"`

	// ID of dataobject
	DataObjectId string `json:"data_object_id"`

	// type of related dataclass
	RelatedDataclassType string `json:"related_dataclass_type"`

	Body *DataobjectSearch `json:"body,omitempty"`
}

func (o ListDataobjectRelationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataobjectRelationRequest struct{}"
	}

	return strings.Join([]string{"ListDataobjectRelationRequest", string(data)}, " ")
}
