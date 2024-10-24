package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CapRef CAP信息
type CapRef struct {

	// CapId，32~36位的英文、数字、短横组合
	CapId *string `json:"cap_id,omitempty"`

	// CapVersionId，32~36位的英文、数字、短横组合
	VersionId *string `json:"version_id,omitempty"`
}

func (o CapRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapRef struct{}"
	}

	return strings.Join([]string{"CapRef", string(data)}, " ")
}
