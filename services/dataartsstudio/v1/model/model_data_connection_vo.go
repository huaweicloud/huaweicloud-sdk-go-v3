package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataConnectionVo struct {

	// 数据连接名称。
	DwName *string `json:"dw_name,omitempty"`

	// 数据连接ID。
	DwId *string `json:"dw_id,omitempty"`

	// 数据连接名称，适配现有实现。
	DisplayName *string `json:"display_name,omitempty"`

	// 数据连接类型。
	DwType *string `json:"dw_type,omitempty"`
}

func (o DataConnectionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataConnectionVo struct{}"
	}

	return strings.Join([]string{"DataConnectionVo", string(data)}, " ")
}
