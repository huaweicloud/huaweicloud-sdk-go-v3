package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateProductRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	Body *ProductCreateRequest `json:"body,omitempty"`
}

func (o CreateProductRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductRequest struct{}"
	}

	return strings.Join([]string{"CreateProductRequest", string(data)}, " ")
}
