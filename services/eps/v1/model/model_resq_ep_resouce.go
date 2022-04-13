package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 获取企业项目下资源请求
type ResqEpResouce struct {
	// 项目ID列表。resource_types中包含region级别服务时为必选项。

	Projects *[]string `json:"projects,omitempty"`
	// 资源类型列表， 此参数为可输入的值（区分大小写）。例如：ecs,scaling_group, images, disk, vpcs,security-groups, shared_bandwidth, eip, cdn等。

	ResourceTypes []string `json:"resource_types"`
	// 索引位置， 从offset指定的下一条数据开始查询，必须为数字，不能为负数，默认为0。

	Offset *int32 `json:"offset,omitempty"`
	// 查询记录数，不传默认为1000，limit最多为1000, 最小值为1。

	Limit *int32 `json:"limit,omitempty"`
	// 搜索字段，key为要匹配的字段，固定为resource_name，value为匹配的值，不传则表示无匹配条件。

	Matches *[]Match `json:"matches,omitempty"`
}

func (o ResqEpResouce) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResqEpResouce struct{}"
	}

	return strings.Join([]string{"ResqEpResouce", string(data)}, " ")
}
