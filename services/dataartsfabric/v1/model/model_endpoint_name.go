package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointName 一个Endpoint名称，只能包含中文、字母、数字、下划线、中划线、点、空格
type EndpointName struct {
}

func (o EndpointName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointName struct{}"
	}

	return strings.Join([]string{"EndpointName", string(data)}, " ")
}
