package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCsrRequest Request Object
type ShowCsrRequest struct {

	// CSR的ID。
	Id string `json:"id"`
}

func (o ShowCsrRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCsrRequest struct{}"
	}

	return strings.Join([]string{"ShowCsrRequest", string(data)}, " ")
}
