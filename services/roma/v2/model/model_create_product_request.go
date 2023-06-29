package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProductRequest Request Object
type CreateProductRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateProductRequestBody `json:"body,omitempty"`
}

func (o CreateProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductRequest struct{}"
	}

	return strings.Join([]string{"CreateProductRequest", string(data)}, " ")
}
