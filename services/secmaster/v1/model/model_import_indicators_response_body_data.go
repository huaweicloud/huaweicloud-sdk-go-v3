package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIndicatorsResponseBodyData 情报导入结果返回数据
type ImportIndicatorsResponseBodyData struct {

	// 导入的情报id列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o ImportIndicatorsResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIndicatorsResponseBodyData struct{}"
	}

	return strings.Join([]string{"ImportIndicatorsResponseBodyData", string(data)}, " ")
}
