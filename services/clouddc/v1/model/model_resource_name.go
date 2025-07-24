package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceName Resource Name Type
type ResourceName struct {
}

func (o ResourceName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceName struct{}"
	}

	return strings.Join([]string{"ResourceName", string(data)}, " ")
}
