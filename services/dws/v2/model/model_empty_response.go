package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EmptyResponse **参数解释**： 请求成功时的空白响应。 **取值范围**： 不涉及。
type EmptyResponse struct {
}

func (o EmptyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EmptyResponse struct{}"
	}

	return strings.Join([]string{"EmptyResponse", string(data)}, " ")
}
