package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowClientRecordResponse Response Object
type AllowClientRecordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AllowClientRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowClientRecordResponse struct{}"
	}

	return strings.Join([]string{"AllowClientRecordResponse", string(data)}, " ")
}
