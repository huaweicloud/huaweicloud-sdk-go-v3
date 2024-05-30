package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchDwByTypeResultData data，统一的返回结果的最外层数据结构。
type SearchDwByTypeResultData struct {

	// 数据连接信息数组
	Value *[]DataConnectionVo `json:"value,omitempty"`
}

func (o SearchDwByTypeResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchDwByTypeResultData struct{}"
	}

	return strings.Join([]string{"SearchDwByTypeResultData", string(data)}, " ")
}
