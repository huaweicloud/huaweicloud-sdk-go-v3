package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoTypeRef struct {
	// 视频类型<br/>

	VideoType *VideoTypeRefVideoType `json:"video_type,omitempty"`
	// 媒资标题</br>

	Title *string `json:"title,omitempty"`
	// 视频描述<br/>

	Description *string `json:"description,omitempty"`
	// 媒资分类id<br/>

	CategoryId *int32 `json:"category_id,omitempty"`
	// 视频标签<br/>

	Tags *string `json:"tags,omitempty"`
	// 是否自动发布<br/>

	AutoPublish *int32 `json:"auto_publish,omitempty"`
	// 转码模板组名称<br/>

	TemplateGroupName *string `json:"template_group_name,omitempty"`
	// 是否自动加密，取值[0，1]，0表示需要不加密；1表示需要加密。加密与转码必须要一起进行，当需要加密时，转码参数不能为空，且转码输出必须要为HLS

	AutoEncrypt *int32 `json:"auto_encrypt,omitempty"`
	// 是否自动预热到CDN,取值[0，1]，0表示不自动预热

	AutoPreheat *int32 `json:"auto_preheat,omitempty"`

	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`

	Review *Review `json:"review,omitempty"`
	// 工作流ID

	WorkflowName *string `json:"workflow_name,omitempty"`
}

func (o VideoTypeRef) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VideoTypeRef struct{}"
	}

	return strings.Join([]string{"VideoTypeRef", string(data)}, " ")
}

type VideoTypeRefVideoType struct {
	value string
}

type VideoTypeRefVideoTypeEnum struct {
	MP4    VideoTypeRefVideoType
	TS     VideoTypeRefVideoType
	MOV    VideoTypeRefVideoType
	MXF    VideoTypeRefVideoType
	MPG    VideoTypeRefVideoType
	FLV    VideoTypeRefVideoType
	WMV    VideoTypeRefVideoType
	HLS    VideoTypeRefVideoType
	DASH   VideoTypeRefVideoType
	MP3    VideoTypeRefVideoType
	WMA    VideoTypeRefVideoType
	APE    VideoTypeRefVideoType
	FLAC   VideoTypeRefVideoType
	AAC    VideoTypeRefVideoType
	AC3    VideoTypeRefVideoType
	MMF    VideoTypeRefVideoType
	AMR    VideoTypeRefVideoType
	M4_A   VideoTypeRefVideoType
	M4_R   VideoTypeRefVideoType
	OGG    VideoTypeRefVideoType
	WAV    VideoTypeRefVideoType
	WV     VideoTypeRefVideoType
	MP2    VideoTypeRefVideoType
	AVI    VideoTypeRefVideoType
	F4_V   VideoTypeRefVideoType
	M4_V   VideoTypeRefVideoType
	MPEG   VideoTypeRefVideoType
	M3_U8  VideoTypeRefVideoType
	E_3_GP VideoTypeRefVideoType
	ASF    VideoTypeRefVideoType
	MKV    VideoTypeRefVideoType
	UNKNOW VideoTypeRefVideoType
}

func GetVideoTypeRefVideoTypeEnum() VideoTypeRefVideoTypeEnum {
	return VideoTypeRefVideoTypeEnum{
		MP4: VideoTypeRefVideoType{
			value: "MP4",
		},
		TS: VideoTypeRefVideoType{
			value: "TS",
		},
		MOV: VideoTypeRefVideoType{
			value: "MOV",
		},
		MXF: VideoTypeRefVideoType{
			value: "MXF",
		},
		MPG: VideoTypeRefVideoType{
			value: "MPG",
		},
		FLV: VideoTypeRefVideoType{
			value: "FLV",
		},
		WMV: VideoTypeRefVideoType{
			value: "WMV",
		},
		HLS: VideoTypeRefVideoType{
			value: "HLS",
		},
		DASH: VideoTypeRefVideoType{
			value: "DASH",
		},
		MP3: VideoTypeRefVideoType{
			value: "MP3",
		},
		WMA: VideoTypeRefVideoType{
			value: "WMA",
		},
		APE: VideoTypeRefVideoType{
			value: "APE",
		},
		FLAC: VideoTypeRefVideoType{
			value: "FLAC",
		},
		AAC: VideoTypeRefVideoType{
			value: "AAC",
		},
		AC3: VideoTypeRefVideoType{
			value: "AC3",
		},
		MMF: VideoTypeRefVideoType{
			value: "MMF",
		},
		AMR: VideoTypeRefVideoType{
			value: "AMR",
		},
		M4_A: VideoTypeRefVideoType{
			value: "M4A",
		},
		M4_R: VideoTypeRefVideoType{
			value: "M4R",
		},
		OGG: VideoTypeRefVideoType{
			value: "OGG",
		},
		WAV: VideoTypeRefVideoType{
			value: "WAV",
		},
		WV: VideoTypeRefVideoType{
			value: "WV",
		},
		MP2: VideoTypeRefVideoType{
			value: "MP2",
		},
		AVI: VideoTypeRefVideoType{
			value: "AVI",
		},
		F4_V: VideoTypeRefVideoType{
			value: "F4V",
		},
		M4_V: VideoTypeRefVideoType{
			value: "M4V",
		},
		MPEG: VideoTypeRefVideoType{
			value: "MPEG",
		},
		M3_U8: VideoTypeRefVideoType{
			value: "M3U8",
		},
		E_3_GP: VideoTypeRefVideoType{
			value: "3GP",
		},
		ASF: VideoTypeRefVideoType{
			value: "ASF",
		},
		MKV: VideoTypeRefVideoType{
			value: "MKV",
		},
		UNKNOW: VideoTypeRefVideoType{
			value: "UNKNOW",
		},
	}
}

func (c VideoTypeRefVideoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *VideoTypeRefVideoType) UnmarshalJSON(b []byte) error {
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
