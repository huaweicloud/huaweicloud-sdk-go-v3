package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCsrRequest Request Object
type DeleteCsrRequest struct {

	// CSR的ID。
	Id string `json:"id"`
}

func (o DeleteCsrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCsrRequest struct{}"
	}

	return strings.Join([]string{"DeleteCsrRequest", string(data)}, " ")
}
