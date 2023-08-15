package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiTestParas struct {

	// page size
	PageSize *string `json:"page_size,omitempty"`

	// page num
	PageNum *string `json:"page_num,omitempty"`
}

func (o ApiTestParas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiTestParas struct{}"
	}

	return strings.Join([]string{"ApiTestParas", string(data)}, " ")
}
