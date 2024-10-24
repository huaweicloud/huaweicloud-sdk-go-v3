package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CapVersionId CapVersionId，32~36位的英文、数字、短横组合
type CapVersionId struct {
}

func (o CapVersionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapVersionId struct{}"
	}

	return strings.Join([]string{"CapVersionId", string(data)}, " ")
}
