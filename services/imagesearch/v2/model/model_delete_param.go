package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteParam struct {

	// 是否幂等删除数据，默认为false。仅对指定ID删除生效。 - false: 数据不存在时返回错误信息。 - true: 数据不存在时返回成功，用于幂等删除场景。
	Force *bool `json:"force,omitempty"`

	// 数据的服务实例级唯一标识，字符长度范围为[1, 256]。 - item_id/custom_tags/custom_num_tags中必须给定至少一个参数，以支持对服务实例中的数据进行指定ID删除或条件删除。 - 如给定item_id参数，则进行指定ID删除，否则进行条件删除。
	ItemId *string `json:"item_id,omitempty"`

	// 自定义字符标签，用于对服务实例中的数据进行custom_tags条件删除。格式为键值对{key:value}。 - key: 必须为服务实例custom_tags中已存在的key，可在创建服务实例时进行配置，或在更新服务实例时进行新增。 - value: 标签值列表，列表内多个标签值为“或”关系，即满足一个即可。列表长度范围为[1, 32]，标签值类型为字符串，字符长度范围为[1, 64]。
	CustomTags map[string][]string `json:"custom_tags,omitempty"`

	// 自定义数值标签，用于对服务实例中的数据进行custom_num_tags条件删除。格式为键值对{key:value}。 - key: 必须为服务实例custom_num_tags中已存在的key，可在创建服务实例时进行配置，或在更新服务实例时进行新增。针对没有设置该数值标签的数据，会直接过滤。 - value: 标签值的取值范围，标签值在给定的取值范围内即视为符合条件。
	CustomNumTags map[string]RangeParam `json:"custom_num_tags,omitempty"`
}

func (o DeleteParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteParam struct{}"
	}

	return strings.Join([]string{"DeleteParam", string(data)}, " ")
}
