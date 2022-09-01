package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LdApiTest struct {

	// 后端API请求参数
	Parameters *[]LdApiParameter `json:"parameters,omitempty" xml:"parameters"`
}

func (o LdApiTest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiTest struct{}"
	}

	return strings.Join([]string{"LdApiTest", string(data)}, " ")
}
