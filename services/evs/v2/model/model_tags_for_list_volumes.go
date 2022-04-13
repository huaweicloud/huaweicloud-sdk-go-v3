package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagsForListVolumes struct {
	// 标签键。

	Key string `json:"key"`
	// 标签值。  标签列表中最多包含10个value。 标签列表中的标签value值不允许重复。 标签列表如果为空列表，表示匹配任意值。标签列表中多个value之间是“或”的关系，在key已经满足要求的前提下，云硬盘满足请求中的某个value就会匹配出来。

	Values []string `json:"values"`
}

func (o TagsForListVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagsForListVolumes struct{}"
	}

	return strings.Join([]string{"TagsForListVolumes", string(data)}, " ")
}
