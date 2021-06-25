package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MetaData struct {
	// 封装类型,TS/MP4等<br/>

	PlayType *int64 `json:"play_type,omitempty"`
	// 封装类型,TS/MP4等<br/>

	PackType *MetaDataPackType `json:"pack_type,omitempty"`
	// 编码类型：H.264、H.265等<br/>

	Codec *MetaDataCodec `json:"codec,omitempty"`
	// 视频时长<br/>

	Duration *int64 `json:"duration,omitempty"`
	// 视频文件大小<br/>

	VideoSize *int64 `json:"video_size,omitempty"`
	// 视频宽度<br/>

	Width *int64 `json:"width,omitempty"`
	// 视频高度<br/>

	Hight *int64 `json:"hight,omitempty"`
	// 视频平均码率<br/>

	BitRate *int64 `json:"bit_rate,omitempty"`
	// 帧率<br/>

	FrameRate *int64 `json:"frame_rate,omitempty"`
	// 分辨率<br/>

	Quality *string `json:"quality,omitempty"`
	// 音频声道数<br/>

	AudioChannels *int32 `json:"audio_channels,omitempty"`
}

func (o MetaData) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MetaData struct{}"
	}

	return strings.Join([]string{"MetaData", string(data)}, " ")
}

type MetaDataPackType struct {
	value string
}

type MetaDataPackTypeEnum struct {
	MP4  MetaDataPackType
	TS   MetaDataPackType
	MOV  MetaDataPackType
	MXF  MetaDataPackType
	MPG  MetaDataPackType
	FLV  MetaDataPackType
	WMV  MetaDataPackType
	MP3  MetaDataPackType
	WMA  MetaDataPackType
	APE  MetaDataPackType
	FLAC MetaDataPackType
	AAC  MetaDataPackType
	AC3  MetaDataPackType
	MMF  MetaDataPackType
	AMR  MetaDataPackType
	M4_A MetaDataPackType
	M4_R MetaDataPackType
	OGG  MetaDataPackType
	WAV  MetaDataPackType
	WV   MetaDataPackType
	MP2  MetaDataPackType
	AVI  MetaDataPackType
	F4_V MetaDataPackType
	M4_V MetaDataPackType
	MPEG MetaDataPackType
	HLS  MetaDataPackType
	DASH MetaDataPackType
}

func GetMetaDataPackTypeEnum() MetaDataPackTypeEnum {
	return MetaDataPackTypeEnum{
		MP4: MetaDataPackType{
			value: "MP4",
		},
		TS: MetaDataPackType{
			value: "TS",
		},
		MOV: MetaDataPackType{
			value: "MOV",
		},
		MXF: MetaDataPackType{
			value: "MXF",
		},
		MPG: MetaDataPackType{
			value: "MPG",
		},
		FLV: MetaDataPackType{
			value: "FLV",
		},
		WMV: MetaDataPackType{
			value: "WMV",
		},
		MP3: MetaDataPackType{
			value: "MP3",
		},
		WMA: MetaDataPackType{
			value: "WMA",
		},
		APE: MetaDataPackType{
			value: "APE",
		},
		FLAC: MetaDataPackType{
			value: "FLAC",
		},
		AAC: MetaDataPackType{
			value: "AAC",
		},
		AC3: MetaDataPackType{
			value: "AC3",
		},
		MMF: MetaDataPackType{
			value: "MMF",
		},
		AMR: MetaDataPackType{
			value: "AMR",
		},
		M4_A: MetaDataPackType{
			value: "M4A",
		},
		M4_R: MetaDataPackType{
			value: "M4R",
		},
		OGG: MetaDataPackType{
			value: "OGG",
		},
		WAV: MetaDataPackType{
			value: "WAV",
		},
		WV: MetaDataPackType{
			value: "WV",
		},
		MP2: MetaDataPackType{
			value: "MP2",
		},
		AVI: MetaDataPackType{
			value: "AVI",
		},
		F4_V: MetaDataPackType{
			value: "F4V",
		},
		M4_V: MetaDataPackType{
			value: "M4V",
		},
		MPEG: MetaDataPackType{
			value: "MPEG",
		},
		HLS: MetaDataPackType{
			value: "HLS",
		},
		DASH: MetaDataPackType{
			value: "DASH",
		},
	}
}

func (c MetaDataPackType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *MetaDataPackType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type MetaDataCodec struct {
	value string
}

type MetaDataCodecEnum struct {
	MPEG_2   MetaDataCodec
	MPEG_4   MetaDataCodec
	H_264    MetaDataCodec
	H_265    MetaDataCodec
	WMV      MetaDataCodec
	VORBIS   MetaDataCodec
	AAC      MetaDataCodec
	EAC_3    MetaDataCodec
	AC_3     MetaDataCodec
	AMR      MetaDataCodec
	APE      MetaDataCodec
	FLAC     MetaDataCodec
	MP3      MetaDataCodec
	MP2      MetaDataCodec
	WMA      MetaDataCodec
	PCM      MetaDataCodec
	ADPCM    MetaDataCodec
	WAV_PACK MetaDataCodec
	HEAAC    MetaDataCodec
	UNKNOWN  MetaDataCodec
}

func GetMetaDataCodecEnum() MetaDataCodecEnum {
	return MetaDataCodecEnum{
		MPEG_2: MetaDataCodec{
			value: "MPEG-2",
		},
		MPEG_4: MetaDataCodec{
			value: "MPEG-4",
		},
		H_264: MetaDataCodec{
			value: "H.264",
		},
		H_265: MetaDataCodec{
			value: "H.265",
		},
		WMV: MetaDataCodec{
			value: "WMV",
		},
		VORBIS: MetaDataCodec{
			value: "Vorbis",
		},
		AAC: MetaDataCodec{
			value: "AAC",
		},
		EAC_3: MetaDataCodec{
			value: "EAC-3",
		},
		AC_3: MetaDataCodec{
			value: "AC-3",
		},
		AMR: MetaDataCodec{
			value: "AMR",
		},
		APE: MetaDataCodec{
			value: "APE",
		},
		FLAC: MetaDataCodec{
			value: "FLAC",
		},
		MP3: MetaDataCodec{
			value: "MP3",
		},
		MP2: MetaDataCodec{
			value: "MP2",
		},
		WMA: MetaDataCodec{
			value: "WMA",
		},
		PCM: MetaDataCodec{
			value: "PCM",
		},
		ADPCM: MetaDataCodec{
			value: "ADPCM",
		},
		WAV_PACK: MetaDataCodec{
			value: "WavPack",
		},
		HEAAC: MetaDataCodec{
			value: "HEAAC",
		},
		UNKNOWN: MetaDataCodec{
			value: "UNKNOWN",
		},
	}
}

func (c MetaDataCodec) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *MetaDataCodec) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
