package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签信息
type TagInfoDto struct {
	// 算子Id

	OperatorId *string `json:"operator_id,omitempty"`
	// 存储ID

	DataStoreId *string `json:"data_store_id,omitempty"`
	// 标签信息

	Tag *string `json:"tag,omitempty"`
}

func (o TagInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagInfoDto struct{}"
	}

	return strings.Join([]string{"TagInfoDto", string(data)}, " ")
}
