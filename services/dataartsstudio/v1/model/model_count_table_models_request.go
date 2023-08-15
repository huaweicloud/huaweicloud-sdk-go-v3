package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTableModelsRequest Request Object
type CountTableModelsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 依据model_id查工作区
	ModelId *int64 `json:"model_id,omitempty"`
}

func (o CountTableModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTableModelsRequest struct{}"
	}

	return strings.Join([]string{"CountTableModelsRequest", string(data)}, " ")
}
