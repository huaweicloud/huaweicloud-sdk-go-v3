package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除事件
type DeleteAlert struct {

	// 删除事件的ID列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o DeleteAlert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlert struct{}"
	}

	return strings.Join([]string{"DeleteAlert", string(data)}, " ")
}
