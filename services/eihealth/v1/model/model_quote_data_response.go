package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuoteDataResponse Response Object
type QuoteDataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o QuoteDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuoteDataResponse struct{}"
	}

	return strings.Join([]string{"QuoteDataResponse", string(data)}, " ")
}
