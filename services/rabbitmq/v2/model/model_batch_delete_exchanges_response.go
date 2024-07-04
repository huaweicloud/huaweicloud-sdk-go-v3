package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteExchangesResponse Response Object
type BatchDeleteExchangesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchDeleteExchangesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteExchangesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteExchangesResponse", string(data)}, " ")
}
