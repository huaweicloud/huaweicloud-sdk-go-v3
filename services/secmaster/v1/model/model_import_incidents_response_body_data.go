package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIncidentsResponseBodyData 事件导入结果返回数据
type ImportIncidentsResponseBodyData struct {

	// 导入的事件id列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o ImportIncidentsResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIncidentsResponseBodyData struct{}"
	}

	return strings.Join([]string{"ImportIncidentsResponseBodyData", string(data)}, " ")
}
