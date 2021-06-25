package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateAssetReq struct {
	// 媒体ID<br/>

	AssetId string `json:"asset_id"`
	// 视频文件MD5值<br/>

	VideoMd5 *string `json:"video_md5,omitempty"`
	// 视频文件名<br/>

	VideoName *string `json:"video_name,omitempty"`
	// 视频文件类型<br/>

	VideoType *UpdateAssetReqVideoType `json:"video_type,omitempty"`
	// 封面ID，取值0-7。当前只支持一张封面，只能填0<br/>

	CoverId *int32 `json:"cover_id,omitempty"`
	// 封面图片格式类型<br/>

	CoverType *UpdateAssetReqCoverType `json:"cover_type,omitempty"`
	// 封面文件MD5值<br/>

	CoverMd5 *string `json:"cover_md5,omitempty"`
	// 字幕文件信息<br/>

	Subtitles *[]Subtitle `json:"subtitles,omitempty"`
}

func (o UpdateAssetReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAssetReq struct{}"
	}

	return strings.Join([]string{"UpdateAssetReq", string(data)}, " ")
}

type UpdateAssetReqVideoType struct {
	value string
}

type UpdateAssetReqVideoTypeEnum struct {
	MP4    UpdateAssetReqVideoType
	TS     UpdateAssetReqVideoType
	MOV    UpdateAssetReqVideoType
	MXF    UpdateAssetReqVideoType
	MPG    UpdateAssetReqVideoType
	FLV    UpdateAssetReqVideoType
	WMV    UpdateAssetReqVideoType
	HLS    UpdateAssetReqVideoType
	MP3    UpdateAssetReqVideoType
	WMA    UpdateAssetReqVideoType
	APE    UpdateAssetReqVideoType
	FLAC   UpdateAssetReqVideoType
	AAC    UpdateAssetReqVideoType
	AC3    UpdateAssetReqVideoType
	MMF    UpdateAssetReqVideoType
	AMR    UpdateAssetReqVideoType
	M4_A   UpdateAssetReqVideoType
	M4_R   UpdateAssetReqVideoType
	OGG    UpdateAssetReqVideoType
	WAV    UpdateAssetReqVideoType
	WV     UpdateAssetReqVideoType
	MP2    UpdateAssetReqVideoType
	AVI    UpdateAssetReqVideoType
	F4_V   UpdateAssetReqVideoType
	M4_V   UpdateAssetReqVideoType
	MPEG   UpdateAssetReqVideoType
	UNKNOW UpdateAssetReqVideoType
}

func GetUpdateAssetReqVideoTypeEnum() UpdateAssetReqVideoTypeEnum {
	return UpdateAssetReqVideoTypeEnum{
		MP4: UpdateAssetReqVideoType{
			value: "MP4",
		},
		TS: UpdateAssetReqVideoType{
			value: "TS",
		},
		MOV: UpdateAssetReqVideoType{
			value: "MOV",
		},
		MXF: UpdateAssetReqVideoType{
			value: "MXF",
		},
		MPG: UpdateAssetReqVideoType{
			value: "MPG",
		},
		FLV: UpdateAssetReqVideoType{
			value: "FLV",
		},
		WMV: UpdateAssetReqVideoType{
			value: "WMV",
		},
		HLS: UpdateAssetReqVideoType{
			value: "HLS",
		},
		MP3: UpdateAssetReqVideoType{
			value: "MP3",
		},
		WMA: UpdateAssetReqVideoType{
			value: "WMA",
		},
		APE: UpdateAssetReqVideoType{
			value: "APE",
		},
		FLAC: UpdateAssetReqVideoType{
			value: "FLAC",
		},
		AAC: UpdateAssetReqVideoType{
			value: "AAC",
		},
		AC3: UpdateAssetReqVideoType{
			value: "AC3",
		},
		MMF: UpdateAssetReqVideoType{
			value: "MMF",
		},
		AMR: UpdateAssetReqVideoType{
			value: "AMR",
		},
		M4_A: UpdateAssetReqVideoType{
			value: "M4A",
		},
		M4_R: UpdateAssetReqVideoType{
			value: "M4R",
		},
		OGG: UpdateAssetReqVideoType{
			value: "OGG",
		},
		WAV: UpdateAssetReqVideoType{
			value: "WAV",
		},
		WV: UpdateAssetReqVideoType{
			value: "WV",
		},
		MP2: UpdateAssetReqVideoType{
			value: "MP2",
		},
		AVI: UpdateAssetReqVideoType{
			value: "AVI",
		},
		F4_V: UpdateAssetReqVideoType{
			value: "F4V",
		},
		M4_V: UpdateAssetReqVideoType{
			value: "M4V",
		},
		MPEG: UpdateAssetReqVideoType{
			value: "MPEG",
		},
		UNKNOW: UpdateAssetReqVideoType{
			value: "UNKNOW",
		},
	}
}

func (c UpdateAssetReqVideoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateAssetReqVideoType) UnmarshalJSON(b []byte) error {
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

type UpdateAssetReqCoverType struct {
	value string
}

type UpdateAssetReqCoverTypeEnum struct {
	JPG UpdateAssetReqCoverType
	PNG UpdateAssetReqCoverType
}

func GetUpdateAssetReqCoverTypeEnum() UpdateAssetReqCoverTypeEnum {
	return UpdateAssetReqCoverTypeEnum{
		JPG: UpdateAssetReqCoverType{
			value: "JPG",
		},
		PNG: UpdateAssetReqCoverType{
			value: "PNG",
		},
	}
}

func (c UpdateAssetReqCoverType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateAssetReqCoverType) UnmarshalJSON(b []byte) error {
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
