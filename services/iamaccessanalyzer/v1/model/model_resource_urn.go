package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceUrn 资源的唯一资源标识符。
type ResourceUrn struct {
}

func (o ResourceUrn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceUrn struct{}"
	}

	return strings.Join([]string{"ResourceUrn", string(data)}, " ")
}
