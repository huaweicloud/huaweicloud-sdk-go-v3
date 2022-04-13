package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateEnterpriseProjectRequest struct {
	Body *EnterpriseProject `json:"body,omitempty"`
}

func (o CreateEnterpriseProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseProjectRequest struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseProjectRequest", string(data)}, " ")
}
