package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListExtendsParamsResponse struct {
	// 附加参数列表

	ExtendsParams *[]ExtendsParamV2 `json:"extends_params,omitempty"`
	// 公共附加参数列表

	CommonParams   *[]CommonParamV2 `json:"common_params,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListExtendsParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExtendsParamsResponse struct{}"
	}

	return strings.Join([]string{"ListExtendsParamsResponse", string(data)}, " ")
}
