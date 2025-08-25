package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateAppsRequest Request Object
type DisassociateAppsRequest struct {
	Body *DisAssociateAppsRequestBody `json:"body,omitempty"`
}

func (o DisassociateAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateAppsRequest struct{}"
	}

	return strings.Join([]string{"DisassociateAppsRequest", string(data)}, " ")
}
