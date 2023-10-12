package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopFactorySupplementDataInstanceResponse Response Object
type StopFactorySupplementDataInstanceResponse struct {

	// success
	Msg *string `json:"msg,omitempty"`

	// 包含若干补数据实例信息
	Rows *[]SupplementDataRespRows `json:"rows,omitempty"`

	// 查询是否成功，取值为true或者false
	Success *bool `json:"success,omitempty"`

	// 补数据实例数量
	Total *int32 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopFactorySupplementDataInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopFactorySupplementDataInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopFactorySupplementDataInstanceResponse", string(data)}, " ")
}
