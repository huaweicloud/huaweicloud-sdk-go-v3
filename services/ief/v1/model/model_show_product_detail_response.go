package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProductDetailResponse Response Object
type ShowProductDetailResponse struct {
	Product        *Product `json:"product,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowProductDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProductDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowProductDetailResponse", string(data)}, " ")
}
