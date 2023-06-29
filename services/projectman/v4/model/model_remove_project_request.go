package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveProjectRequest Request Object
type RemoveProjectRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`
}

func (o RemoveProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveProjectRequest struct{}"
	}

	return strings.Join([]string{"RemoveProjectRequest", string(data)}, " ")
}
