package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDomainRequest Request Object
type DeleteDomainRequest struct {
	Body *DeleteDomainV2RequestBody `json:"body,omitempty"`
}

func (o DeleteDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainRequest struct{}"
	}

	return strings.Join([]string{"DeleteDomainRequest", string(data)}, " ")
}
