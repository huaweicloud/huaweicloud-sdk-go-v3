package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectDomainRequest Request Object
type UpdateProjectDomainRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 领域id
	DomainId string `json:"domain_id"`

	Body *CreateProjectDomainRequestBody `json:"body,omitempty"`
}

func (o UpdateProjectDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectDomainRequest struct{}"
	}

	return strings.Join([]string{"UpdateProjectDomainRequest", string(data)}, " ")
}
