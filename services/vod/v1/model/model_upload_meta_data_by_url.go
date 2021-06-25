package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UploadMetaDataByUrl struct {
	// 视频类型<br/>

	VideoType *UploadMetaDataByUrlVideoType `json:"video_type,omitempty"`
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
	// 视频源文件URL</br>

	Url *string `json:"url,omitempty"`
}

func (o UploadMetaDataByUrl) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UploadMetaDataByUrl struct{}"
	}

	return strings.Join([]string{"UploadMetaDataByUrl", string(data)}, " ")
}

type UploadMetaDataByUrlVideoType struct {
	value string
}

type UploadMetaDataByUrlVideoTypeEnum struct {
	MP4    UploadMetaDataByUrlVideoType
	TS     UploadMetaDataByUrlVideoType
	MOV    UploadMetaDataByUrlVideoType
	MXF    UploadMetaDataByUrlVideoType
	MPG    UploadMetaDataByUrlVideoType
	FLV    UploadMetaDataByUrlVideoType
	WMV    UploadMetaDataByUrlVideoType
	HLS    UploadMetaDataByUrlVideoType
	DASH   UploadMetaDataByUrlVideoType
	MP3    UploadMetaDataByUrlVideoType
	WMA    UploadMetaDataByUrlVideoType
	APE    UploadMetaDataByUrlVideoType
	FLAC   UploadMetaDataByUrlVideoType
	AAC    UploadMetaDataByUrlVideoType
	AC3    UploadMetaDataByUrlVideoType
	MMF    UploadMetaDataByUrlVideoType
	AMR    UploadMetaDataByUrlVideoType
	M4_A   UploadMetaDataByUrlVideoType
	M4_R   UploadMetaDataByUrlVideoType
	OGG    UploadMetaDataByUrlVideoType
	WAV    UploadMetaDataByUrlVideoType
	WV     UploadMetaDataByUrlVideoType
	MP2    UploadMetaDataByUrlVideoType
	AVI    UploadMetaDataByUrlVideoType
	F4_V   UploadMetaDataByUrlVideoType
	M4_V   UploadMetaDataByUrlVideoType
	MPEG   UploadMetaDataByUrlVideoType
	M3_U8  UploadMetaDataByUrlVideoType
	E_3_GP UploadMetaDataByUrlVideoType
	ASF    UploadMetaDataByUrlVideoType
	MKV    UploadMetaDataByUrlVideoType
	UNKNOW UploadMetaDataByUrlVideoType
}

func GetUploadMetaDataByUrlVideoTypeEnum() UploadMetaDataByUrlVideoTypeEnum {
	return UploadMetaDataByUrlVideoTypeEnum{
		MP4: UploadMetaDataByUrlVideoType{
			value: "MP4",
		},
		TS: UploadMetaDataByUrlVideoType{
			value: "TS",
		},
		MOV: UploadMetaDataByUrlVideoType{
			value: "MOV",
		},
		MXF: UploadMetaDataByUrlVideoType{
			value: "MXF",
		},
		MPG: UploadMetaDataByUrlVideoType{
			value: "MPG",
		},
		FLV: UploadMetaDataByUrlVideoType{
			value: "FLV",
		},
		WMV: UploadMetaDataByUrlVideoType{
			value: "WMV",
		},
		HLS: UploadMetaDataByUrlVideoType{
			value: "HLS",
		},
		DASH: UploadMetaDataByUrlVideoType{
			value: "DASH",
		},
		MP3: UploadMetaDataByUrlVideoType{
			value: "MP3",
		},
		WMA: UploadMetaDataByUrlVideoType{
			value: "WMA",
		},
		APE: UploadMetaDataByUrlVideoType{
			value: "APE",
		},
		FLAC: UploadMetaDataByUrlVideoType{
			value: "FLAC",
		},
		AAC: UploadMetaDataByUrlVideoType{
			value: "AAC",
		},
		AC3: UploadMetaDataByUrlVideoType{
			value: "AC3",
		},
		MMF: UploadMetaDataByUrlVideoType{
			value: "MMF",
		},
		AMR: UploadMetaDataByUrlVideoType{
			value: "AMR",
		},
		M4_A: UploadMetaDataByUrlVideoType{
			value: "M4A",
		},
		M4_R: UploadMetaDataByUrlVideoType{
			value: "M4R",
		},
		OGG: UploadMetaDataByUrlVideoType{
			value: "OGG",
		},
		WAV: UploadMetaDataByUrlVideoType{
			value: "WAV",
		},
		WV: UploadMetaDataByUrlVideoType{
			value: "WV",
		},
		MP2: UploadMetaDataByUrlVideoType{
			value: "MP2",
		},
		AVI: UploadMetaDataByUrlVideoType{
			value: "AVI",
		},
		F4_V: UploadMetaDataByUrlVideoType{
			value: "F4V",
		},
		M4_V: UploadMetaDataByUrlVideoType{
			value: "M4V",
		},
		MPEG: UploadMetaDataByUrlVideoType{
			value: "MPEG",
		},
		M3_U8: UploadMetaDataByUrlVideoType{
			value: "M3U8",
		},
		E_3_GP: UploadMetaDataByUrlVideoType{
			value: "3GP",
		},
		ASF: UploadMetaDataByUrlVideoType{
			value: "ASF",
		},
		MKV: UploadMetaDataByUrlVideoType{
			value: "MKV",
		},
		UNKNOW: UploadMetaDataByUrlVideoType{
			value: "UNKNOW",
		},
	}
}

func (c UploadMetaDataByUrlVideoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UploadMetaDataByUrlVideoType) UnmarshalJSON(b []byte) error {
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
