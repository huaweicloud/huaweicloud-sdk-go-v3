package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输出为dis类型时的配置信息
type TaskOutputDis struct {
	// DIS流名称，选用dis类型输出时必填。

	StreamName string `json:"stream_name"`
	// 作业输出数据类别的列表，选填，仅部分服务需要。当输出类型下有这个列表时，表示希望这个输出类型下存放dataCategory列表内的数据。

	DataCategory *[]string `json:"data_category,omitempty"`
}

func (o TaskOutputDis) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputDis struct{}"
	}

	return strings.Join([]string{"TaskOutputDis", string(data)}, " ")
}
