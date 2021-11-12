package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 输出为hosting类型时的配置信息
type TaskOutputHosting struct {
	// 作业输出数据类别的列表，选填，仅部分服务需要。当输出类型下有这个列表时，表示希望这个输出类型下存放dataCategory列表内的数据。

	DataCategory *[]string `json:"data_category,omitempty"`
}

func (o TaskOutputHosting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskOutputHosting struct{}"
	}

	return strings.Join([]string{"TaskOutputHosting", string(data)}, " ")
}
