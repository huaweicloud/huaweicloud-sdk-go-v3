package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Remarks 备注
type Remarks struct {
}

func (o Remarks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Remarks struct{}"
	}

	return strings.Join([]string{"Remarks", string(data)}, " ")
}
