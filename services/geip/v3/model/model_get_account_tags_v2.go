package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAccountTagsV2 资源标签。
type GetAccountTagsV2 struct {

	// - 功能说明：标签名称 - 约束：   - 创建的预定义标签如果与已有的预定义标签完全相同，则会覆盖已有的预定义标签；若只有“键”相同，“值”不同，则为新创建的预定义标签。   - 键的长度最大36字符，由英文字母、数字、下划线、中划线、中文字符组成。   - 单个资源最多可以添加20个标签。
	Key string `json:"key"`

	// - 功能说明：标签值列表 - 约束：   - 值的长度最大43字符，由英文字母、数字、下划线、点、中划线、中文字符组成。
	Values []string `json:"values"`
}

func (o GetAccountTagsV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAccountTagsV2 struct{}"
	}

	return strings.Join([]string{"GetAccountTagsV2", string(data)}, " ")
}
