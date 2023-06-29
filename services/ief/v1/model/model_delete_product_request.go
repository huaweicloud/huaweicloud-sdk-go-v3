package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteProductRequest Request Object
type DeleteProductRequest struct {

	// 批量节点注册作业ID
	ProductId string `json:"product_id"`

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`
}

func (o DeleteProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProductRequest struct{}"
	}

	return strings.Join([]string{"DeleteProductRequest", string(data)}, " ")
}
