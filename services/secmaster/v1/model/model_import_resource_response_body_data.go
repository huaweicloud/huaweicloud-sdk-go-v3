package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportResourceResponseBodyData 告警导入结果返回数据
type ImportResourceResponseBodyData struct {

	// 导入的告警id列表
	BatchIds *[]string `json:"batch_ids,omitempty"`
}

func (o ImportResourceResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportResourceResponseBodyData struct{}"
	}

	return strings.Join([]string{"ImportResourceResponseBodyData", string(data)}, " ")
}
