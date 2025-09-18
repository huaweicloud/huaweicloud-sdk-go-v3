package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoCommon struct {

	// 输出策略。  取值如下： - discard - transcode  >- 当视频参数中的“output_policy”为\"discard\"，且音频参数中的“output_policy”为“transcode”时，表示只输出音频。 >- 当视频参数中的“output_policy”为\"transcode\"，且音频参数中的“output_policy”为“discard”时，表示只输出视频。 >- 同时为\"discard\"时不合法。 >- 同时为“transcode”时，表示输出音视频。
	OutputPolicy *VideoCommonOutputPolicy `json:"output_policy,omitempty"`

	// 视频编码格式。  取值如下： - 1：表示H.264。 - 2：表示H.265。
	Codec *int32 `json:"codec,omitempty"`

	// 编码档次，建议设为3。  取值如下： - 1：VIDEO_PROFILE_H264_BASE - 2：VIDEO_PROFILE_H264_MAIN - 3：VIDEO_PROFILE_H264_HIGH - 4：VIDEO_PROFILE_H265_MAIN
	Profile *int32 `json:"profile,omitempty"`

	// 编码级别。  取值如下： - 1：VIDEO_LEVEL_1_0 - 2：VIDEO_LEVEL_1_1 - 3：VIDEO_LEVEL_1_2 - 4：VIDEO_LEVEL_1_3 - 5：VIDEO_LEVEL_2_0 - 6：VIDEO_LEVEL_2_1 - 7：VIDEO_LEVEL_2_2 - 8：VIDEO_LEVEL_3_0 - 9：VIDEO_LEVEL_3_1 - 10：VIDEO_LEVEL_3_2 - 11：VIDEO_LEVEL_4_0 - 12：VIDEO_LEVEL_4_1 - 13：VIDEO_LEVEL_4_2 - 14：VIDEO_LEVEL_5_0 - 15：VIDEO_LEVEL_5_1
	Level *int32 `json:"level,omitempty"`

	// 编码质量等级。  取值如下： - 1：VIDEO_PRESET_SPEED，编码快速档位 - 3：VIDEO_PRESET_HIGHQUALITY，编码高质量档位 - 4：VIDEO_PRESET_QUALITY，编码质量档位 - 5：VIDEO_PRESET_BALANCE，编码平衡档位 默认值1。
	Preset *int32 `json:"preset,omitempty"`

	// I帧最大间隔  取值范围：[2，10]。  默认值：5。  单位：秒。
	MaxIframesInterval *int32 `json:"max_iframes_interval,omitempty"`

	// 最大B帧间隔。  取值范围： - H264：[0，7]，默认值为4。 - H265：[0，7]，默认值为7。  单位：帧。
	BframesCount *int32 `json:"bframes_count,omitempty"`

	// 帧率  取值范围：0或[5,60]之间的整数，0表示自适应  单位：帧每秒
	FrameRate *int32 `json:"frame_rate,omitempty"`

	// 黑边剪裁类型  取值如下： - 0：不开启黑边剪裁 - 1：开启黑边剪裁，低复杂度算法，针对长视频（>5分钟） - 2：开启黑边剪裁，高复杂度算法，针对短视频（<=5分钟）
	BlackCut *int32 `json:"black_cut,omitempty"`
}

func (o VideoCommon) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoCommon struct{}"
	}

	return strings.Join([]string{"VideoCommon", string(data)}, " ")
}

type VideoCommonOutputPolicy struct {
	value string
}

type VideoCommonOutputPolicyEnum struct {
	TRANSCODE VideoCommonOutputPolicy
	DISCARD   VideoCommonOutputPolicy
	COPY      VideoCommonOutputPolicy
}

func GetVideoCommonOutputPolicyEnum() VideoCommonOutputPolicyEnum {
	return VideoCommonOutputPolicyEnum{
		TRANSCODE: VideoCommonOutputPolicy{
			value: "transcode",
		},
		DISCARD: VideoCommonOutputPolicy{
			value: "discard",
		},
		COPY: VideoCommonOutputPolicy{
			value: "copy",
		},
	}
}

func (c VideoCommonOutputPolicy) Value() string {
	return c.value
}

func (c VideoCommonOutputPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoCommonOutputPolicy) UnmarshalJSON(b []byte) error {
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
