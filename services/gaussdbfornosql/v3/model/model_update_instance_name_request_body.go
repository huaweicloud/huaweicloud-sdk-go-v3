package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceNameRequestBody struct {

	// 新实例名称。用于表示实例的名称。取值范围：长度为4~64位，必须以字母开头（A~Z或a~z），区分大小写，可以包含字母、数字（0~9）、中划线（-）或者下划线（_），不能包含其他特殊字符。
	Name string `json:"name"`
}

func (o UpdateInstanceNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceNameRequestBody", string(data)}, " ")
}
