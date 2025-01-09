package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCsrRequestBody struct {

	// CSR名称。
	Name string `json:"name"`
}

func (o UpdateCsrRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCsrRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCsrRequestBody", string(data)}, " ")
}
