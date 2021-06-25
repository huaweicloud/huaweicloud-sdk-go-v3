package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateAssetByFileReq struct {
	// 视频类型<br/>

	VideoType *CreateAssetByFileReqVideoType `json:"video_type,omitempty"`
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
	// 视频文件名<br/>

	VideoName *string `json:"video_name,omitempty"`
	// 视频文件MD5值<br/>

	VideoMd5 *string `json:"video_md5,omitempty"`
	// 封面图片文件类型<br/>

	CoverType *CreateAssetByFileReqCoverType `json:"cover_type,omitempty"`
	// 封面文件MD5值<br/>

	CoverMd5 *string `json:"cover_md5,omitempty"`
	// 字幕文件信息<br/>

	Subtitles *[]Subtitle `json:"subtitles,omitempty"`
}

func (o CreateAssetByFileReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAssetByFileReq struct{}"
	}

	return strings.Join([]string{"CreateAssetByFileReq", string(data)}, " ")
}

type CreateAssetByFileReqVideoType struct {
	value string
}

type CreateAssetByFileReqVideoTypeEnum struct {
	MP4    CreateAssetByFileReqVideoType
	TS     CreateAssetByFileReqVideoType
	MOV    CreateAssetByFileReqVideoType
	MXF    CreateAssetByFileReqVideoType
	MPG    CreateAssetByFileReqVideoType
	FLV    CreateAssetByFileReqVideoType
	WMV    CreateAssetByFileReqVideoType
	HLS    CreateAssetByFileReqVideoType
	DASH   CreateAssetByFileReqVideoType
	MP3    CreateAssetByFileReqVideoType
	WMA    CreateAssetByFileReqVideoType
	APE    CreateAssetByFileReqVideoType
	FLAC   CreateAssetByFileReqVideoType
	AAC    CreateAssetByFileReqVideoType
	AC3    CreateAssetByFileReqVideoType
	MMF    CreateAssetByFileReqVideoType
	AMR    CreateAssetByFileReqVideoType
	M4_A   CreateAssetByFileReqVideoType
	M4_R   CreateAssetByFileReqVideoType
	OGG    CreateAssetByFileReqVideoType
	WAV    CreateAssetByFileReqVideoType
	WV     CreateAssetByFileReqVideoType
	MP2    CreateAssetByFileReqVideoType
	AVI    CreateAssetByFileReqVideoType
	F4_V   CreateAssetByFileReqVideoType
	M4_V   CreateAssetByFileReqVideoType
	MPEG   CreateAssetByFileReqVideoType
	M3_U8  CreateAssetByFileReqVideoType
	E_3_GP CreateAssetByFileReqVideoType
	ASF    CreateAssetByFileReqVideoType
	MKV    CreateAssetByFileReqVideoType
	UNKNOW CreateAssetByFileReqVideoType
}

func GetCreateAssetByFileReqVideoTypeEnum() CreateAssetByFileReqVideoTypeEnum {
	return CreateAssetByFileReqVideoTypeEnum{
		MP4: CreateAssetByFileReqVideoType{
			value: "MP4",
		},
		TS: CreateAssetByFileReqVideoType{
			value: "TS",
		},
		MOV: CreateAssetByFileReqVideoType{
			value: "MOV",
		},
		MXF: CreateAssetByFileReqVideoType{
			value: "MXF",
		},
		MPG: CreateAssetByFileReqVideoType{
			value: "MPG",
		},
		FLV: CreateAssetByFileReqVideoType{
			value: "FLV",
		},
		WMV: CreateAssetByFileReqVideoType{
			value: "WMV",
		},
		HLS: CreateAssetByFileReqVideoType{
			value: "HLS",
		},
		DASH: CreateAssetByFileReqVideoType{
			value: "DASH",
		},
		MP3: CreateAssetByFileReqVideoType{
			value: "MP3",
		},
		WMA: CreateAssetByFileReqVideoType{
			value: "WMA",
		},
		APE: CreateAssetByFileReqVideoType{
			value: "APE",
		},
		FLAC: CreateAssetByFileReqVideoType{
			value: "FLAC",
		},
		AAC: CreateAssetByFileReqVideoType{
			value: "AAC",
		},
		AC3: CreateAssetByFileReqVideoType{
			value: "AC3",
		},
		MMF: CreateAssetByFileReqVideoType{
			value: "MMF",
		},
		AMR: CreateAssetByFileReqVideoType{
			value: "AMR",
		},
		M4_A: CreateAssetByFileReqVideoType{
			value: "M4A",
		},
		M4_R: CreateAssetByFileReqVideoType{
			value: "M4R",
		},
		OGG: CreateAssetByFileReqVideoType{
			value: "OGG",
		},
		WAV: CreateAssetByFileReqVideoType{
			value: "WAV",
		},
		WV: CreateAssetByFileReqVideoType{
			value: "WV",
		},
		MP2: CreateAssetByFileReqVideoType{
			value: "MP2",
		},
		AVI: CreateAssetByFileReqVideoType{
			value: "AVI",
		},
		F4_V: CreateAssetByFileReqVideoType{
			value: "F4V",
		},
		M4_V: CreateAssetByFileReqVideoType{
			value: "M4V",
		},
		MPEG: CreateAssetByFileReqVideoType{
			value: "MPEG",
		},
		M3_U8: CreateAssetByFileReqVideoType{
			value: "M3U8",
		},
		E_3_GP: CreateAssetByFileReqVideoType{
			value: "3GP",
		},
		ASF: CreateAssetByFileReqVideoType{
			value: "ASF",
		},
		MKV: CreateAssetByFileReqVideoType{
			value: "MKV",
		},
		UNKNOW: CreateAssetByFileReqVideoType{
			value: "UNKNOW",
		},
	}
}

func (c CreateAssetByFileReqVideoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateAssetByFileReqVideoType) UnmarshalJSON(b []byte) error {
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

type CreateAssetByFileReqCoverType struct {
	value string
}

type CreateAssetByFileReqCoverTypeEnum struct {
	JPG CreateAssetByFileReqCoverType
	PNG CreateAssetByFileReqCoverType
}

func GetCreateAssetByFileReqCoverTypeEnum() CreateAssetByFileReqCoverTypeEnum {
	return CreateAssetByFileReqCoverTypeEnum{
		JPG: CreateAssetByFileReqCoverType{
			value: "JPG",
		},
		PNG: CreateAssetByFileReqCoverType{
			value: "PNG",
		},
	}
}

func (c CreateAssetByFileReqCoverType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateAssetByFileReqCoverType) UnmarshalJSON(b []byte) error {
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
