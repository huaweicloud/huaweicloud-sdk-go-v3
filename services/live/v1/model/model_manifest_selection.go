package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ManifestSelection 流选择配置
type ManifestSelection struct {

	// **参数解释**： 流排序模式 **约束限制**： 不涉及 **取值范围**： - ORIGINAL：保持原始顺序，即按照频道配置中转码模板的顺序 - VIDEO_BITRATE_ASCENDING：按照视频码率升序 - VIDEO_BITRATE_DESCENDING：按照视频码率降序
	StreamOrder *ManifestSelectionStreamOrder `json:"stream_order,omitempty"`

	// **参数解释**： 视频码率过滤的最小码率 **约束限制**： 单位：bit/s；取值必须比max_video_bandwidth小
	MinVideoBandwidth *int32 `json:"min_video_bandwidth,omitempty"`

	// **参数解释**： 视频码率过滤的最大码率 **约束限制**： 单位：bit/s；取值必须比min_video_bandwidth大
	MaxVideoBandwidth *int32 `json:"max_video_bandwidth,omitempty"`
}

func (o ManifestSelection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ManifestSelection struct{}"
	}

	return strings.Join([]string{"ManifestSelection", string(data)}, " ")
}

type ManifestSelectionStreamOrder struct {
	value string
}

type ManifestSelectionStreamOrderEnum struct {
	ORIGINAL                 ManifestSelectionStreamOrder
	VIDEO_BITRATE_ASCENDING  ManifestSelectionStreamOrder
	VIDEO_BITRATE_DESCENDING ManifestSelectionStreamOrder
}

func GetManifestSelectionStreamOrderEnum() ManifestSelectionStreamOrderEnum {
	return ManifestSelectionStreamOrderEnum{
		ORIGINAL: ManifestSelectionStreamOrder{
			value: "ORIGINAL",
		},
		VIDEO_BITRATE_ASCENDING: ManifestSelectionStreamOrder{
			value: "VIDEO_BITRATE_ASCENDING",
		},
		VIDEO_BITRATE_DESCENDING: ManifestSelectionStreamOrder{
			value: "VIDEO_BITRATE_DESCENDING",
		},
	}
}

func (c ManifestSelectionStreamOrder) Value() string {
	return c.value
}

func (c ManifestSelectionStreamOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ManifestSelectionStreamOrder) UnmarshalJSON(b []byte) error {
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
