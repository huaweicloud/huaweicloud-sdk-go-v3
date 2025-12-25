package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTagValueRequestBody 更新标签值时参数对象
type UpdateTagValueRequestBody struct {

	// 原标签值
	OldValue string `json:"old_value"`

	// 新标签值
	NewValue string `json:"new_value"`
}

func (o UpdateTagValueRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTagValueRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTagValueRequestBody", string(data)}, " ")
}
