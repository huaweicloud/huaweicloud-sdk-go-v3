package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListConnectionNumberDataList struct {

	// 时间戳毫秒值
	Time *int32 `json:"time,omitempty"`

	// 连接数
	Value *int32 `json:"value,omitempty"`
}

func (o ListConnectionNumberDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionNumberDataList struct{}"
	}

	return strings.Join([]string{"ListConnectionNumberDataList", string(data)}, " ")
}
