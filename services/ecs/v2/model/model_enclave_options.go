package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnclaveOptions struct {
	Enabled *bool `json:"enabled,omitempty"`
}

func (o EnclaveOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnclaveOptions struct{}"
	}

	return strings.Join([]string{"EnclaveOptions", string(data)}, " ")
}
