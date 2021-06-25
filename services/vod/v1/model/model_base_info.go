package model

import (
	"encoding/json"

	"strings"
)

type BaseInfo struct {
	// 媒资状态。 \"CREATING\"   //上传中 \"FAILED\"     //上传失败 \"CREATED\"  //上传成功 \"PUBLISHED\"  //已发布 \"DELETED\"  //已删除

	AssetStatus *string `json:"asset_status,omitempty"`
	// 媒体标题<br/>

	Title *string `json:"title,omitempty"`
	// 视频文件名<br/>

	VideoName *string `json:"video_name,omitempty"`
	// 视频描述<br/>

	Description *string `json:"description,omitempty"`
	// 媒资分类id<br/>

	CategoryId *int64 `json:"category_id,omitempty"`
	// 媒资分类名称<br/>

	CategoryName *string `json:"category_name,omitempty"`
	// 媒资创建时间<br/>

	CreateTime *string `json:"create_time,omitempty"`
	// 媒资最近修改时间<br/>

	LastModified *string `json:"last_modified,omitempty"`
	// 音视频文件类型。 取值如下： 视频文件：MP4、TS、MOV、MXF、MPG、FLV、WMV、AVI、M4V、F4V、MPEG、3GP、ASF、MKV 音频文件：MP3、OGG、WAV、WMA、APE、FLAC、AAC、AC3、MMF、AMR、M4A、M4R、WV、MP2

	VideoType *string `json:"video_type,omitempty"`
	// 视频标签<br/>

	Tags *string `json:"tags,omitempty"`

	MetaData *MetaData `json:"meta_data,omitempty"`
	// 原始视频文件的下载地址<br/>

	VideoUrl *string `json:"video_url,omitempty"`
	// 封面信息数组<br/>

	CoverInfoArray *[]CoverInfo `json:"cover_info_array,omitempty"`
	// 字幕信息数组<br/>

	SubtitleInfo *[]SubtitleInfo `json:"subtitle_info,omitempty"`

	SourchPath *FileAddr `json:"sourch_path,omitempty"`

	SourcePath *FileAddr `json:"source_path,omitempty"`

	OutputPath *FileAddr `json:"output_path,omitempty"`
	// 创建该资源的用户<br/>

	CreateUser *string `json:"create_user,omitempty"`
	// 权限<br/>

	Privilege *string `json:"privilege,omitempty"`
	// 文件夹code<br/>

	FolderName *string `json:"folder_name,omitempty"`
	// 自定义元数据<br/>

	CustomMetadata map[string]interface{} `json:"custom_metadata,omitempty"`
}

func (o BaseInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BaseInfo struct{}"
	}

	return strings.Join([]string{"BaseInfo", string(data)}, " ")
}
