package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHyperinstanceGroupRequest Request Object
type CreateHyperinstanceGroupRequest struct {
	Body *HyperinstanceGroupCreateRequest `json:"body,omitempty"`
}

func (o CreateHyperinstanceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHyperinstanceGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateHyperinstanceGroupRequest", string(data)}, " ")
}
