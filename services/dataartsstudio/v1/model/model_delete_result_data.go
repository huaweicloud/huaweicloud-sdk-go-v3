package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResultData 删除的最终返回结果，返回成功删除的对象个数。
type DeleteResultData struct {

	// 成功删除的对象个数
	Value *string `json:"value,omitempty"`
}

func (o DeleteResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResultData struct{}"
	}

	return strings.Join([]string{"DeleteResultData", string(data)}, " ")
}
