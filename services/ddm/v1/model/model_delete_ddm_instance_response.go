package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDdmInstanceResponse Response Object
type DeleteDdmInstanceResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteDdmInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDdmInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDdmInstanceResponse", string(data)}, " ")
}
