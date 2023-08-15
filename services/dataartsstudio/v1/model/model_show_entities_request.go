package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEntitiesRequest Request Object
type ShowEntitiesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *OpenEntitySearchRequest `json:"body,omitempty"`
}

func (o ShowEntitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEntitiesRequest struct{}"
	}

	return strings.Join([]string{"ShowEntitiesRequest", string(data)}, " ")
}
