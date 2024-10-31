package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpReferenceTableResponse Response Object
type DeleteHttpReferenceTableResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteHttpReferenceTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpReferenceTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteHttpReferenceTableResponse", string(data)}, " ")
}
