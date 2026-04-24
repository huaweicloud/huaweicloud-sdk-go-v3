package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainNewRequest Request Object
type UpdateDomainNewRequest struct {
	Body *UpdateDomainNewReq `json:"body,omitempty"`
}

func (o UpdateDomainNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainNewRequest struct{}"
	}

	return strings.Join([]string{"UpdateDomainNewRequest", string(data)}, " ")
}
