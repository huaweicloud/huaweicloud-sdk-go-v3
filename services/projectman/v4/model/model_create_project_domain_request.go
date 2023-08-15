package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectDomainRequest Request Object
type CreateProjectDomainRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *CreateProjectDomainRequestBody `json:"body,omitempty"`
}

func (o CreateProjectDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectDomainRequest struct{}"
	}

	return strings.Join([]string{"CreateProjectDomainRequest", string(data)}, " ")
}
