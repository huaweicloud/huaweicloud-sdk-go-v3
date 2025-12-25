package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CuQuotaAmount 数量
type CuQuotaAmount struct {
}

func (o CuQuotaAmount) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CuQuotaAmount struct{}"
	}

	return strings.Join([]string{"CuQuotaAmount", string(data)}, " ")
}
