package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceCode 服务名称缩写。
type ServiceCode struct {
}

func (o ServiceCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceCode struct{}"
	}

	return strings.Join([]string{"ServiceCode", string(data)}, " ")
}
