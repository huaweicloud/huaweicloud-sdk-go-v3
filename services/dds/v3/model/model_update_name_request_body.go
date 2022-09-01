package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNameRequestBody struct {

	// 新实例名称。用于表示实例的名称，允许和已有名称重复。取值范围：长度为4~64位，必须以字母开头（A~Z或a~z），区分大小写，可以包含字母、数字（0~9）、中划线（-）或者下划线（_），不能包含其他特殊字符。
	NewInstanceName string `json:"new_instance_name" xml:"new_instance_name"`
}

func (o UpdateNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNameRequestBody", string(data)}, " ")
}
