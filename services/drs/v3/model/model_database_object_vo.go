package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatabaseObjectVo 数据库对象信息体
type DatabaseObjectVo struct {

	// 数据库对象和数据库表名称，例如格式为hec-*-*-drs_test1。
	Id *string `json:"id,omitempty"`

	// 是否选择高级配置，值为true。
	Select *string `json:"select,omitempty"`
}

func (o DatabaseObjectVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObjectVo struct{}"
	}

	return strings.Join([]string{"DatabaseObjectVo", string(data)}, " ")
}
