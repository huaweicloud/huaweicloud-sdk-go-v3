package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountAllModelsRequest Request Object
type CountAllModelsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`
}

func (o CountAllModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountAllModelsRequest struct{}"
	}

	return strings.Join([]string{"CountAllModelsRequest", string(data)}, " ")
}
