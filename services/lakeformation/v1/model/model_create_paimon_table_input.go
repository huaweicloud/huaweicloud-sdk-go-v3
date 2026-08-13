package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePaimonTableInput Paimon表格式结构体
type CreatePaimonTableInput struct {
	Schema *PaimonSchema `json:"schema"`
}

func (o CreatePaimonTableInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePaimonTableInput struct{}"
	}

	return strings.Join([]string{"CreatePaimonTableInput", string(data)}, " ")
}
