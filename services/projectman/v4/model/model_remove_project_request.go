package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RemoveProjectRequest struct {
	// devcloud的项目id

	ProjectId string `json:"project_id"`
}

func (o RemoveProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveProjectRequest struct{}"
	}

	return strings.Join([]string{"RemoveProjectRequest", string(data)}, " ")
}
