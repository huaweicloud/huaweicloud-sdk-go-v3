package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatasetInfo dataset input when grant policy
type DatasetInfo struct {

	// 数据集名称。只能包含中文、字母、数字和_*-特殊字符，且长度为1~256。
	Name string `json:"name"`
}

func (o DatasetInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasetInfo struct{}"
	}

	return strings.Join([]string{"DatasetInfo", string(data)}, " ")
}
