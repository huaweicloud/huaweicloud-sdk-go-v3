package model

import (
	"encoding/json"

	"strings"
)

// 灾备初始化进度
type StructProcessResp struct {
	// 数据生成时间

	CreateTime *string `json:"create_time,omitempty"`
	// 对比结果

	Result *[]StructProcessVo `json:"result,omitempty"`
}

func (o StructProcessResp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StructProcessResp struct{}"
	}

	return strings.Join([]string{"StructProcessResp", string(data)}, " ")
}
