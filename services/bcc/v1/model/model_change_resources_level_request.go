package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeResourcesLevelRequest Request Object
type ChangeResourcesLevelRequest struct {

	// 账号ID
	DomainId string `json:"domain_id"`

	Body *UpdateResourceLevelBody `json:"body,omitempty"`
}

func (o ChangeResourcesLevelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeResourcesLevelRequest struct{}"
	}

	return strings.Join([]string{"ChangeResourcesLevelRequest", string(data)}, " ")
}
