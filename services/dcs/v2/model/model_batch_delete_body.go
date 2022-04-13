package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteBody struct {
	// 实例的ID列表。  仅当URI中参数all_failure值为“false”或者其他值时，才需要配置该参数。

	Instances *[]string `json:"instances,omitempty"`
}

func (o BatchDeleteBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteBody", string(data)}, " ")
}
