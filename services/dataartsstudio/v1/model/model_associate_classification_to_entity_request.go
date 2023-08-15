package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateClassificationToEntityRequest Request Object
type AssociateClassificationToEntityRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 资产标识guid
	Guid string `json:"guid"`

	Body *OpenClassification `json:"body,omitempty"`
}

func (o AssociateClassificationToEntityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateClassificationToEntityRequest struct{}"
	}

	return strings.Join([]string{"AssociateClassificationToEntityRequest", string(data)}, " ")
}
