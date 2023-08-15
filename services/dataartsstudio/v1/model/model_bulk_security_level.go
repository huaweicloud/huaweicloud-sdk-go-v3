package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BulkSecurityLevel struct {

	// 资产guid
	Guids []string `json:"guids"`

	// 密级
	SecurityLevel string `json:"security_level"`
}

func (o BulkSecurityLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BulkSecurityLevel struct{}"
	}

	return strings.Join([]string{"BulkSecurityLevel", string(data)}, " ")
}
