package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProductRequest Request Object
type UpdateProductRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 产品ID
	ProductId int32 `json:"product_id"`

	Body *UpdateProductRequestBody `json:"body,omitempty"`
}

func (o UpdateProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductRequest struct{}"
	}

	return strings.Join([]string{"UpdateProductRequest", string(data)}, " ")
}
