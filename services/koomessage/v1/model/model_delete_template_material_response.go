package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTemplateMaterialResponse Response Object
type DeleteTemplateMaterialResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 响应消息。
	Message *string `json:"message,omitempty"`

	// 固定为null
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteTemplateMaterialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTemplateMaterialResponse struct{}"
	}

	return strings.Join([]string{"DeleteTemplateMaterialResponse", string(data)}, " ")
}
