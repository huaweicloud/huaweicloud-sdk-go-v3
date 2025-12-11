package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserStatisticInfoResponseInfo 账号统计信息
type UserStatisticInfoResponseInfo struct {

	// **参数解释**: 资产中的账号名称，用于唯一标识账号 **取值范围**: 字符长度1-64位，参考Windows文件命名规则，支持字母、数字、下划线、中文及!@.-等特殊字符，不含中文标点符号
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 当前账号关联的资产（主机/容器）数量 **取值范围**: 非负整数，最小值0；单位：台（主机）/个（容器）
	Num *int32 `json:"num,omitempty"`
}

func (o UserStatisticInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserStatisticInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"UserStatisticInfoResponseInfo", string(data)}, " ")
}
