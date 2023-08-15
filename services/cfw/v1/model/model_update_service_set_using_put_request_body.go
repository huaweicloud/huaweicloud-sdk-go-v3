package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServiceSetUsingPutRequestBody Update Service Set Required Body
type UpdateServiceSetUsingPutRequestBody struct {

	// 服务组名称
	Name *string `json:"name,omitempty"`

	// 服务组描述信息
	Description *string `json:"description,omitempty"`
}

func (o UpdateServiceSetUsingPutRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceSetUsingPutRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateServiceSetUsingPutRequestBody", string(data)}, " ")
}
