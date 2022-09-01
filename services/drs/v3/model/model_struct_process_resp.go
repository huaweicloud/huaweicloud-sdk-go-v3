package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 灾备初始化进度
type StructProcessResp struct {

	// 数据生成时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 对比结果
	Result *[]StructProcessVo `json:"result,omitempty" xml:"result"`
}

func (o StructProcessResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StructProcessResp struct{}"
	}

	return strings.Join([]string{"StructProcessResp", string(data)}, " ")
}
