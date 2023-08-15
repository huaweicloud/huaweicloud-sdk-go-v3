package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProductRequest Request Object
type DeleteProductRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 产品ID
	ProductId int32 `json:"product_id"`
}

func (o DeleteProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProductRequest struct{}"
	}

	return strings.Join([]string{"DeleteProductRequest", string(data)}, " ")
}
