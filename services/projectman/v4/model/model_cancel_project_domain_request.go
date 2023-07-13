package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelProjectDomainRequest Request Object
type CancelProjectDomainRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 领域id
	DomainId string `json:"domain_id"`
}

func (o CancelProjectDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelProjectDomainRequest struct{}"
	}

	return strings.Join([]string{"CancelProjectDomainRequest", string(data)}, " ")
}
