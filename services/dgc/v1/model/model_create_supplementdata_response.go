package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSupplementdataResponse Response Object
type CreateSupplementdataResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSupplementdataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSupplementdataResponse struct{}"
	}

	return strings.Join([]string{"CreateSupplementdataResponse", string(data)}, " ")
}
