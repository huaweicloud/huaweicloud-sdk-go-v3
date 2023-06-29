package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnterpriseProjectRequest Request Object
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
