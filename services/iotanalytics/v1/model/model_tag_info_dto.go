package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签信息
type TagInfoDto struct {

	// 算子Id
	OperatorId *string `json:"operator_id,omitempty" xml:"operator_id"`

	// 存储ID
	DataStoreId *string `json:"data_store_id,omitempty" xml:"data_store_id"`

	// 标签信息
	Tag *string `json:"tag,omitempty" xml:"tag"`
}

func (o TagInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagInfoDto struct{}"
	}

	return strings.Join([]string{"TagInfoDto", string(data)}, " ")
}
