package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityUserTablePermissionRequest Request Object
type ListSecurityUserTablePermissionRequest struct {

	// dataarts实例id
	Instance *string `json:"instance,omitempty"`

	Body *SecurityListUserTableList `json:"body,omitempty"`
}

func (o ListSecurityUserTablePermissionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityUserTablePermissionRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityUserTablePermissionRequest", string(data)}, " ")
}
