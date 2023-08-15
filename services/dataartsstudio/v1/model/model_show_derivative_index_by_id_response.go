package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDerivativeIndexByIdResponse Response Object
type ShowDerivativeIndexByIdResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDerivativeIndexByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDerivativeIndexByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowDerivativeIndexByIdResponse", string(data)}, " ")
}
