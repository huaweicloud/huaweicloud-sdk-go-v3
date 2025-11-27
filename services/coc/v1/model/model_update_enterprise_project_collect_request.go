package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnterpriseProjectCollectRequest Request Object
type UpdateEnterpriseProjectCollectRequest struct {
	Body *UpdateEnterpriseProjectCollectRequestBody `json:"body,omitempty"`
}

func (o UpdateEnterpriseProjectCollectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseProjectCollectRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseProjectCollectRequest", string(data)}, " ")
}
