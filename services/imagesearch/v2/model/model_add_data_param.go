package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddDataParam struct {

	// 是否强制添加数据，默认为false。 - false: 数据已存在则不进行添加。 - true: 数据已存在仍然覆盖添加。
	Force *bool `json:"force,omitempty"`

	// 数据的服务实例级唯一标识，字符长度范围为[1, 256]。
	ItemId string `json:"item_id"`

	// 数据的描述信息，字符长度范围为[1, 2048]。
	Desc *string `json:"desc,omitempty"`

	// 数据的自定义字符标签，可用于搜索时的过滤。格式为键值对{key:value}。 - key: 必须为服务实例custom_tags中已存在的key，可在创建服务实例时进行配置，或在更新服务实例时进行新增。 - value: 类型为字符串，字符长度范围为[1, 64]。
	CustomTags map[string]string `json:"custom_tags,omitempty"`

	// 数据的自定义数值标签，可用于搜索时的过滤。格式为键值对{key:value}。 - key: 必须为服务实例custom_num_tags中已存在的key，可在创建服务实例时进行配置，或在更新服务实例时进行新增。 - value: 类型为数值，格式为double。
	CustomNumTags map[string]float64 `json:"custom_num_tags,omitempty"`

	// 图像文件的base64字符串，图像入库时，与image_url二选一。要求如下： - 格式：目前仅支持JPEG/JPG/PNG/BMP格式的图像。 - 尺寸：默认情况下，要求图像的最短边大于128px，最长边小于4096px。部分服务类型有特殊要求，可参见服务类型说明。 - 大小：图像文件大小要求不超过5M。 - 其他：图片中不能包含旋转信息。
	ImageBase64 *string `json:"image_base64,omitempty"`

	// 图像文件的服务可访问URL，字符长度范围为[1, 4096]。图像入库时，与image_base64二选一。
	ImageUrl *string `json:"image_url,omitempty"`

	// 关键词列表，关键词数量范围为[1, 100]，关键词字符长度范围为[1, 64]。支持关键词搜索的服务实例添加数据时，如给定此参数，则直接使用给定的keywords作为关键词，否则会自动生成对应的keywords。
	Keywords *[]string `json:"keywords,omitempty"`

	OptionalParams *AddDataOptionalParam `json:"optional_params,omitempty"`
}

func (o AddDataParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDataParam struct{}"
	}

	return strings.Join([]string{"AddDataParam", string(data)}, " ")
}
