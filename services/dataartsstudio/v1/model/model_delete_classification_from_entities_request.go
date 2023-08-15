package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClassificationFromEntitiesRequest Request Object
type DeleteClassificationFromEntitiesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 资产标识guid
	Guid string `json:"guid"`

	Body *OpenClassification `json:"body,omitempty"`
}

func (o DeleteClassificationFromEntitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClassificationFromEntitiesRequest struct{}"
	}

	return strings.Join([]string{"DeleteClassificationFromEntitiesRequest", string(data)}, " ")
}
