package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TotalCount 资源总数量。
type TotalCount struct {
}

func (o TotalCount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TotalCount struct{}"
	}

	return strings.Join([]string{"TotalCount", string(data)}, " ")
}
