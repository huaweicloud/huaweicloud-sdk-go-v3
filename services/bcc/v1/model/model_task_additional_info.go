package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TaskAdditionalInfo struct {

	// 附加信息条数
	ItemNumber int32 `json:"item_number"`

	// 附加信息
	AdditionalInfo map[string]string `json:"additional_info"`
}

func (o TaskAdditionalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskAdditionalInfo struct{}"
	}

	return strings.Join([]string{"TaskAdditionalInfo", string(data)}, " ")
}
