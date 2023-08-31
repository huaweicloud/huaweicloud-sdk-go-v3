package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppResponse Response Object
type CreateAppResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	Result         *AppBaseResponse `json:"result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppResponse struct{}"
	}

	return strings.Join([]string{"CreateAppResponse", string(data)}, " ")
}
