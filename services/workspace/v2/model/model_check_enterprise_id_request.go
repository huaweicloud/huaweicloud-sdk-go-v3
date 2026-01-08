package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckEnterpriseIdRequest Request Object
type CheckEnterpriseIdRequest struct {
	Body *CheckEnterpriseIdReq `json:"body,omitempty"`
}

func (o CheckEnterpriseIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckEnterpriseIdRequest struct{}"
	}

	return strings.Join([]string{"CheckEnterpriseIdRequest", string(data)}, " ")
}
