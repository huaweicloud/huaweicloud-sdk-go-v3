package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceNamespaceRequest Request Object
type CreateInstanceNamespaceRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateInstanceNamespaceRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceNamespaceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceNamespaceRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceNamespaceRequest", string(data)}, " ")
}
