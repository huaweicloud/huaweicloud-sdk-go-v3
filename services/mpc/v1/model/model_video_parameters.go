package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoParameters struct {

	// 输出策略。  取值如下： - discard - transcode  >- 当视频参数中的“output_policy”为\"discard\"，且音频参数中的“output_policy”为“transcode”时，表示只输出音频。 >- 当视频参数中的“output_policy”为\"transcode\"，且音频参数中的“output_policy”为“discard”时，表示只输出视频。 >- 同时为\"discard\"时不合法。 >- 同时为“transcode”时，表示输出音视频。
	OutputPolicy *VideoParametersOutputPolicy `json:"output_policy,omitempty"`

	// 视频编码格式。  取值如下： - 1：VIDEO_CODEC_H264 - 2：VIDEO_CODEC_H265
	Codec *int32 `json:"codec,omitempty"`

	// 视频恒定码率控制因子。  取值范围为[0, 51]
	Crf *interface{} `json:"crf,omitempty"`

	// 输出最大码率，基于crf，设置max_bitrate字段才会开启ccrf  取值范围：0或[40,800000]之间的整数。   单位：kbit/s  带crf时使用，参考原片的平均码率进行设置（一般为1.5倍）
	MaxBitrate *int32 `json:"max_bitrate,omitempty"`

	// 输出平均码率。  取值范围：0或[40,50000]之间的整数。  单位：kbit/s  若设置为0，则输出平均码率为自适应值。
	Bitrate *int32 `json:"bitrate,omitempty"`

	// ccrf时的缓冲区大小,建议与max_bitrate保持一致，避免编码器缓冲区溢出  取值范围：0或[40,800000]之间的整数。  单位：kbit
	BufSize *int32 `json:"buf_size,omitempty"`

	// 编码档次  取值如下： - 1：VIDEO_PROFILE_H264_BASE - 2：VIDEO_PROFILE_H264_MAIN - 3：VIDEO_PROFILE_H264_HIGH - 4：VIDEO_PROFILE_H265_MAIN
	Profile *int32 `json:"profile,omitempty"`

	// 编码级别  取值如下： - 1：VIDEO_LEVEL_1_0 - 2：VIDEO_LEVEL_1_1 - 3：VIDEO_LEVEL_1_2 - 4：VIDEO_LEVEL_1_3 - 5：VIDEO_LEVEL_2_0 - 6：VIDEO_LEVEL_2_1 - 7：VIDEO_LEVEL_2_2 - 8：VIDEO_LEVEL_3_0 - 9：VIDEO_LEVEL_3_1 - 10：VIDEO_LEVEL_3_2 - 11：VIDEO_LEVEL_4_0 - 12：VIDEO_LEVEL_4_1 - 13：VIDEO_LEVEL_4_2 - 14：VIDEO_LEVEL_5_0 - 15：VIDEO_LEVEL_5_1 - 16：VIDEO_LEVEL_x_x
	Level *int32 `json:"level,omitempty"`

	// 编码质量等级  取值如下： - 1：VIDEO_PRESET_SPEED，编码快速档位 - 3：VIDEO_PRESET_HIGHQUALITY，编码高质量档位 - 4：VIDEO_PRESET_QUALITY，编码质量档位 - 5：VIDEO_PRESET_BALANCE，编码平衡档位  默认值1。
	Preset *int32 `json:"preset,omitempty"`

	// I帧最大间隔  取值范围：[2，10]。  默认值：5。  单位：秒。
	MaxIframesInterval *int32 `json:"max_iframes_interval,omitempty"`

	// 最大B帧间隔。  取值范围： - H264：[0，7]，默认值为4。 - H265：[0，7]，默认值为7。  单位：帧。
	BframesCount *int32 `json:"bframes_count,omitempty"`

	// 帧率。  取值范围：0或[5,60]，0表示自适应。  单位：帧每秒。  > 若设置的帧率不在取值范围内，则自动调整为0，若设置的帧率高于片源帧率，则自动调整为片源帧率。
	FrameRate *int32 `json:"frame_rate,omitempty"`

	// 帧率。  取值范围：0或[5,60]，0表示自适应。  单位：帧每秒。  > 若设置的帧率不在取值范围内，则自动调整为0，若设置的帧率高于片源帧率，则自动调整为片源帧率。
	FrameRateFloat *float32 `json:"frame_rate_float,omitempty"`

	// 视频宽度（单位：像素）  - H264：范围[32,4096]，必须为2的倍数 - H265：范围[320,4096]，必须是4的倍数
	Width *int32 `json:"width,omitempty"`

	// 视频高度（单位：像素）  - H264：范围[32,2880]，必须为2的倍数 - H265：范围[240,2880] ，必须是4的倍数
	Height *int32 `json:"height,omitempty"`

	// 黑边剪裁类型  - 0：不开启黑边剪裁 - 1：开启黑边剪裁，低复杂度算法，针对长视频（>5分钟） - 2：开启黑边剪裁，高复杂度算法，针对短视频（<=5分钟）
	BlackCut *int32 `json:"black_cut,omitempty"`

	// 流名称
	StreamName *string `json:"stream_name,omitempty"`
}

func (o VideoParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoParameters struct{}"
	}

	return strings.Join([]string{"VideoParameters", string(data)}, " ")
}

type VideoParametersOutputPolicy struct {
	value string
}

type VideoParametersOutputPolicyEnum struct {
	TRANSCODE VideoParametersOutputPolicy
	DISCARD   VideoParametersOutputPolicy
	COPY      VideoParametersOutputPolicy
}

func GetVideoParametersOutputPolicyEnum() VideoParametersOutputPolicyEnum {
	return VideoParametersOutputPolicyEnum{
		TRANSCODE: VideoParametersOutputPolicy{
			value: "transcode",
		},
		DISCARD: VideoParametersOutputPolicy{
			value: "discard",
		},
		COPY: VideoParametersOutputPolicy{
			value: "copy",
		},
	}
}

func (c VideoParametersOutputPolicy) Value() string {
	return c.value
}

func (c VideoParametersOutputPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VideoParametersOutputPolicy) UnmarshalJSON(b []byte) error {
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
