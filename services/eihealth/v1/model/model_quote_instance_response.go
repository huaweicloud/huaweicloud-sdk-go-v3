package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuoteInstanceResponse Response Object
type QuoteInstanceResponse struct {
	Body           *[]QuoteDatabaseResultRsp `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o QuoteInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuoteInstanceResponse struct{}"
	}

	return strings.Join([]string{"QuoteInstanceResponse", string(data)}, " ")
}
