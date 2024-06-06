package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SearchParam struct {

	// 搜索类型，必须为服务实例支持的搜索类型。服务实例的搜索类型列表可在创建服务实例时进行配置。 > 可以使用枚举名或者枚举值（例如IMAGE/0），枚举值可能会变动，建议使用枚举名。
	SearchType SearchParamSearchType `json:"search_type"`

	// 返回搜索结果的数量，默认为10，取值范围为[1, 1000]。
	Limit *int32 `json:"limit,omitempty"`

	// 返回搜索结果的偏移量，即返回序号在[offset, offset+limit]内的搜索结果。默认为0，取值范围为[0, N]。 - 默认情况下，搜索要求offset+limit <= 1000。 - 针对支持全量召回的场景，使用全量召回时，要求offset必须为0。
	Offset *int32 `json:"offset,omitempty"`

	LastItem *SearchAfterParam `json:"last_item,omitempty"`

	// 返回搜索结果的最小得分，用于对搜索结果进行score过滤，取值范围为[0, 1]。 - 目前仅对IMAGE/CATEGORY搜索类型生效。
	MinScore *float64 `json:"min_score,omitempty"`

	// 自定义字符标签，用于对搜索结果进行条件过滤。格式为键值对{key:value}。 - key: 必须为服务实例custom_tags中已存在的key，可在创建服务实例时进行配置，或在更新服务实例时进行新增。 - value: 标签值列表，列表内多个标签值为“或”关系，即满足一个即可。列表长度范围为[1, 32]，标签值类型为字符串，字符长度范围为[1, 64]。
	CustomTags map[string][]string `json:"custom_tags,omitempty"`

	// 自定义数值标签，用于对搜索结果进行条件过滤。格式为键值对{key:value}。 - key: 必须为服务实例custom_num_tags中已存在的key，可在创建服务实例时进行配置，或在更新服务实例时进行新增。针对没有设置该数值标签的数据，会直接过滤。 - value: 标签值的取值范围，标签值在给定的取值范围内即视为符合条件。
	CustomNumTags map[string]RangeParam `json:"custom_num_tags,omitempty"`

	// 图像文件的base64字符串，基于图像搜索时，与image_url二选一。要求如下： - 格式：目前仅支持JPEG/JPG/PNG/BMP/WEBP格式的图像。 - 大小：图像文件大小要求不超过5M。 - 尺寸：默认情况下，要求图像的最短边大于64px，最长边小于4096px。部分服务类型有特殊要求，可参见服务类型说明。 - 其他：图片中不能包含旋转信息。
	ImageBase64 *string `json:"image_base64,omitempty"`

	// 图像文件的服务可访问URL，字符长度范围为[1, 4096]。基于图像搜索时，与image_base64二选一。
	ImageUrl *string `json:"image_url,omitempty"`

	// 关键词列表，搜索时关键词数量范围为[1, 10]，关键词字符长度范围为[1, 64]。使用KEYWORD搜索类型进行搜索时，必须提供该参数。
	Keywords *[]string `json:"keywords,omitempty"`

	// 文本字符串，字符长度范围为[1, 512]。
	Text *string `json:"text,omitempty"`

	OptionalParams *SearchOptionalParam `json:"optional_params,omitempty"`
}

func (o SearchParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchParam struct{}"
	}

	return strings.Join([]string{"SearchParam", string(data)}, " ")
}

type SearchParamSearchType struct {
	value string
}

type SearchParamSearchTypeEnum struct {
	IMAGE    SearchParamSearchType
	KEYWORD  SearchParamSearchType
	TEXT     SearchParamSearchType
	CATEGORY SearchParamSearchType
}

func GetSearchParamSearchTypeEnum() SearchParamSearchTypeEnum {
	return SearchParamSearchTypeEnum{
		IMAGE: SearchParamSearchType{
			value: "IMAGE",
		},
		KEYWORD: SearchParamSearchType{
			value: "KEYWORD",
		},
		TEXT: SearchParamSearchType{
			value: "TEXT",
		},
		CATEGORY: SearchParamSearchType{
			value: "CATEGORY",
		},
	}
}

func (c SearchParamSearchType) Value() string {
	return c.value
}

func (c SearchParamSearchType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SearchParamSearchType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
