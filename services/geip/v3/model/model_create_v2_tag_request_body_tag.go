package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateV2TagRequestBodyTag 标签对象
type CreateV2TagRequestBodyTag struct {

	// - 功能说明：标签名称 - 约束：   - 创建的预定义标签如果与已有的预定义标签完全相同，则会覆盖已有的预定义标签；若只有“键”相同，“值”不同，则为新创建的预定义标签。   - 键的长度最大36字符，由英文字母、数字、下划线、中划线、中文字符组成。   - 单个资源最多可以添加20个标签。
	Key string `json:"key"`

	// - 功能说明：标签值 - 约束：   - 值的长度最大43字符，由英文字母、数字、下划线、点、中划线、中文字符组成。
	Value string `json:"value"`
}

func (o CreateV2TagRequestBodyTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2TagRequestBodyTag struct{}"
	}

	return strings.Join([]string{"CreateV2TagRequestBodyTag", string(data)}, " ")
}
