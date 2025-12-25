package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportAlertResponseBodyData 告警导入结果返回数据
type ImportAlertResponseBodyData struct {

	// 导入的告警id列表
	Ids *[]string `json:"ids,omitempty"`
}

func (o ImportAlertResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAlertResponseBodyData struct{}"
	}

	return strings.Join([]string{"ImportAlertResponseBodyData", string(data)}, " ")
}
