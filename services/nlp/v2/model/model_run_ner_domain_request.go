package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunNerDomainRequest struct {
	Body *PostDomainNerRequest `json:"body,omitempty"`
}

func (o RunNerDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunNerDomainRequest struct{}"
	}

	return strings.Join([]string{"RunNerDomainRequest", string(data)}, " ")
}
