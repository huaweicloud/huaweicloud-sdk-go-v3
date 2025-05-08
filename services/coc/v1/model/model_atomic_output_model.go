package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AtomicOutputModel struct {
	Supported *bool `json:"supported,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o AtomicOutputModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtomicOutputModel struct{}"
	}

	return strings.Join([]string{"AtomicOutputModel", string(data)}, " ")
}
