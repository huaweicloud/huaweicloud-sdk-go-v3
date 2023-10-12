package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExecutionParamsResponse Response Object
type ShowExecutionParamsResponse struct {

	// 查询部署记录执行参数返回体
	Body           *[]ConfigInfo `json:"body,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowExecutionParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExecutionParamsResponse struct{}"
	}

	return strings.Join([]string{"ShowExecutionParamsResponse", string(data)}, " ")
}
