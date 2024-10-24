package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceDescription 描述信息
type ResourceDescription struct {
}

func (o ResourceDescription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDescription struct{}"
	}

	return strings.Join([]string{"ResourceDescription", string(data)}, " ")
}
