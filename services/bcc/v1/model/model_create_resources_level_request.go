package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourcesLevelRequest Request Object
type CreateResourcesLevelRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	Body *CreateResourceLevelBody `json:"body,omitempty"`
}

func (o CreateResourcesLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourcesLevelRequest struct{}"
	}

	return strings.Join([]string{"CreateResourcesLevelRequest", string(data)}, " ")
}
