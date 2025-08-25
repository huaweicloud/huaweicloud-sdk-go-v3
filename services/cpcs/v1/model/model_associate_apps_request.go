package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateAppsRequest Request Object
type AssociateAppsRequest struct {
	Body *AssociateAppsRequestBody `json:"body,omitempty"`
}

func (o AssociateAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAppsRequest struct{}"
	}

	return strings.Join([]string{"AssociateAppsRequest", string(data)}, " ")
}
