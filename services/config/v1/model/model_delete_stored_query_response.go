package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStoredQueryResponse Response Object
type DeleteStoredQueryResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStoredQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStoredQueryResponse struct{}"
	}

	return strings.Join([]string{"DeleteStoredQueryResponse", string(data)}, " ")
}
