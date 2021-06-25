package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PublishAssetFromObsReq struct {
	// 视频类型<br/>

	VideoType *PublishAssetFromObsReqVideoType `json:"video_type,omitempty"`
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
	// 存储模式：0表示视频拷贝到点播桶 1表示视频存在租户的桶<br/>

	StorageMode *int32 `json:"storage_mode,omitempty"`

	Input *FileAddr `json:"input,omitempty"`
	// 输出桶，storage_mode为1的时候才选择此参数 <br/>

	OutputBucket *string `json:"output_bucket,omitempty"`
	// 输出路径，storage_mode为1的时候才选择此参数<br/>

	OutputPath *string `json:"output_path,omitempty"`
}

func (o PublishAssetFromObsReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PublishAssetFromObsReq struct{}"
	}

	return strings.Join([]string{"PublishAssetFromObsReq", string(data)}, " ")
}

type PublishAssetFromObsReqVideoType struct {
	value string
}

type PublishAssetFromObsReqVideoTypeEnum struct {
	MP4    PublishAssetFromObsReqVideoType
	TS     PublishAssetFromObsReqVideoType
	MOV    PublishAssetFromObsReqVideoType
	MXF    PublishAssetFromObsReqVideoType
	MPG    PublishAssetFromObsReqVideoType
	FLV    PublishAssetFromObsReqVideoType
	WMV    PublishAssetFromObsReqVideoType
	HLS    PublishAssetFromObsReqVideoType
	DASH   PublishAssetFromObsReqVideoType
	MP3    PublishAssetFromObsReqVideoType
	WMA    PublishAssetFromObsReqVideoType
	APE    PublishAssetFromObsReqVideoType
	FLAC   PublishAssetFromObsReqVideoType
	AAC    PublishAssetFromObsReqVideoType
	AC3    PublishAssetFromObsReqVideoType
	MMF    PublishAssetFromObsReqVideoType
	AMR    PublishAssetFromObsReqVideoType
	M4_A   PublishAssetFromObsReqVideoType
	M4_R   PublishAssetFromObsReqVideoType
	OGG    PublishAssetFromObsReqVideoType
	WAV    PublishAssetFromObsReqVideoType
	WV     PublishAssetFromObsReqVideoType
	MP2    PublishAssetFromObsReqVideoType
	AVI    PublishAssetFromObsReqVideoType
	F4_V   PublishAssetFromObsReqVideoType
	M4_V   PublishAssetFromObsReqVideoType
	MPEG   PublishAssetFromObsReqVideoType
	M3_U8  PublishAssetFromObsReqVideoType
	E_3_GP PublishAssetFromObsReqVideoType
	ASF    PublishAssetFromObsReqVideoType
	MKV    PublishAssetFromObsReqVideoType
	UNKNOW PublishAssetFromObsReqVideoType
}

func GetPublishAssetFromObsReqVideoTypeEnum() PublishAssetFromObsReqVideoTypeEnum {
	return PublishAssetFromObsReqVideoTypeEnum{
		MP4: PublishAssetFromObsReqVideoType{
			value: "MP4",
		},
		TS: PublishAssetFromObsReqVideoType{
			value: "TS",
		},
		MOV: PublishAssetFromObsReqVideoType{
			value: "MOV",
		},
		MXF: PublishAssetFromObsReqVideoType{
			value: "MXF",
		},
		MPG: PublishAssetFromObsReqVideoType{
			value: "MPG",
		},
		FLV: PublishAssetFromObsReqVideoType{
			value: "FLV",
		},
		WMV: PublishAssetFromObsReqVideoType{
			value: "WMV",
		},
		HLS: PublishAssetFromObsReqVideoType{
			value: "HLS",
		},
		DASH: PublishAssetFromObsReqVideoType{
			value: "DASH",
		},
		MP3: PublishAssetFromObsReqVideoType{
			value: "MP3",
		},
		WMA: PublishAssetFromObsReqVideoType{
			value: "WMA",
		},
		APE: PublishAssetFromObsReqVideoType{
			value: "APE",
		},
		FLAC: PublishAssetFromObsReqVideoType{
			value: "FLAC",
		},
		AAC: PublishAssetFromObsReqVideoType{
			value: "AAC",
		},
		AC3: PublishAssetFromObsReqVideoType{
			value: "AC3",
		},
		MMF: PublishAssetFromObsReqVideoType{
			value: "MMF",
		},
		AMR: PublishAssetFromObsReqVideoType{
			value: "AMR",
		},
		M4_A: PublishAssetFromObsReqVideoType{
			value: "M4A",
		},
		M4_R: PublishAssetFromObsReqVideoType{
			value: "M4R",
		},
		OGG: PublishAssetFromObsReqVideoType{
			value: "OGG",
		},
		WAV: PublishAssetFromObsReqVideoType{
			value: "WAV",
		},
		WV: PublishAssetFromObsReqVideoType{
			value: "WV",
		},
		MP2: PublishAssetFromObsReqVideoType{
			value: "MP2",
		},
		AVI: PublishAssetFromObsReqVideoType{
			value: "AVI",
		},
		F4_V: PublishAssetFromObsReqVideoType{
			value: "F4V",
		},
		M4_V: PublishAssetFromObsReqVideoType{
			value: "M4V",
		},
		MPEG: PublishAssetFromObsReqVideoType{
			value: "MPEG",
		},
		M3_U8: PublishAssetFromObsReqVideoType{
			value: "M3U8",
		},
		E_3_GP: PublishAssetFromObsReqVideoType{
			value: "3GP",
		},
		ASF: PublishAssetFromObsReqVideoType{
			value: "ASF",
		},
		MKV: PublishAssetFromObsReqVideoType{
			value: "MKV",
		},
		UNKNOW: PublishAssetFromObsReqVideoType{
			value: "UNKNOW",
		},
	}
}

func (c PublishAssetFromObsReqVideoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *PublishAssetFromObsReqVideoType) UnmarshalJSON(b []byte) error {
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
