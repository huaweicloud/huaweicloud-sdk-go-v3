package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteCustomizedCheckitemsRequestBody struct {

	// 删除自定义检查项的id列表
	BatchIds []string `json:"batch_ids"`
}

func (o DeleteCustomizedCheckitemsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomizedCheckitemsRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteCustomizedCheckitemsRequestBody", string(data)}, " ")
}
