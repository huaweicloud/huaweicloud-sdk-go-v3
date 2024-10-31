package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceInstanceName 一个Service Instance的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
type ServiceInstanceName struct {
}

func (o ServiceInstanceName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceInstanceName struct{}"
	}

	return strings.Join([]string{"ServiceInstanceName", string(data)}, " ")
}
