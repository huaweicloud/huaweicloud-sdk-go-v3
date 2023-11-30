package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDerivativeIndexByIdResponse Response Object
type ShowDerivativeIndexByIdResponse struct {
	Data           *DerivativeIndexVoDetailData `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowDerivativeIndexByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDerivativeIndexByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowDerivativeIndexByIdResponse", string(data)}, " ")
}
