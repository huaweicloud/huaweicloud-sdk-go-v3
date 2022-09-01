package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProjectModuleRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id" xml:"project_id"`

	Body *CreateProjectModuleRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CreateProjectModuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectModuleRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectModuleRequest", string(data)}, " ")
}