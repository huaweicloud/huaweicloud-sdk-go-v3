package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FrameDecodeConfig 泛协议码流的拆包组包配置
type FrameDecodeConfig struct {

	// **参数说明**：拆包组包规则。 **取值范围**： - DELIMITER：通过特定分隔符（如逗号、换行符等）来拆分或组合数据包。 - FIXED_LENGTH：按照固定的字节长度，对每一帧数据进行拆分或组合。 - FIELD_LENGTH：每一帧的长度可变，通过数据包中携带的长度字段信息进行拆分或组合。
	FrameDecodeType *string `json:"frame_decode_type,omitempty"`

	// **参数说明**：单个帧的最大长度。拆包规则为DELIMITER|FIELD_LENGTH时，该参数必选。
	MaxFrameLength *int32 `json:"max_frame_length,omitempty"`

	// **参数说明**：分隔符，hex string格式。拆包规则为DELIMITER，该参数必选。
	Delimiter *string `json:"delimiter,omitempty"`

	// **参数说明**：单个帧的固定长度。拆包规则为FIXED_LENGTH，该参数必选。
	FixedFrameLength *int32 `json:"fixed_frame_length,omitempty"`

	// **参数说明**：指定长度字段在数据包中的起始位置（偏移量）。拆包规则为FIELD_LENGTH ，该参数必选。
	FieldOffset *int32 `json:"field_offset,omitempty"`

	// **参数说明**：指定长度字段占用的字节数。拆包规则为FIELD_LENGTH，该参数必选。
	FieldLength *int32 `json:"field_length,omitempty"`

	// **参数说明**：起始字符，hex string格式。拆包规则为FIXED_LENGTH，该参数可选。
	InitialBytes *string `json:"initial_bytes,omitempty"`

	// **参数说明**：调整长度字段的值。拆包规则为FIELD_LENGTH，该参数可选。
	AdjustmentLength *int32 `json:"adjustment_length,omitempty"`

	// **参数说明**：指定解码后从数据包中去掉的字节数。通常用于去掉长度字段，只保留数据内容。拆包规则为FIELD_LENGTH，该参数可选。
	InitialBytesToStrip *int32 `json:"initial_bytes_to_strip,omitempty"`
}

func (o FrameDecodeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrameDecodeConfig struct{}"
	}

	return strings.Join([]string{"FrameDecodeConfig", string(data)}, " ")
}
