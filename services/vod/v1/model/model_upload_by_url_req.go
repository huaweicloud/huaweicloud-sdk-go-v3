package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadByUrlReq struct {

	// OBS的bucket名称。
	Bucket *string `json:"bucket,omitempty"`

	// 文件的存储路径。
	Object *string `json:"object,omitempty"`

	// 视频源文件URL</br>
	Url *string `json:"url,omitempty"`

	// 视频类型<br/>
	FileType *string `json:"file_type,omitempty"`

	// 媒资标题</br>
	Title *string `json:"title,omitempty"`

	// 回调url
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 自定义上下文，回调时会回调给用户，透传信息
	SessionContext *string `json:"session_context,omitempty"`
}

func (o UploadByUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadByUrlReq struct{}"
	}

	return strings.Join([]string{"UploadByUrlReq", string(data)}, " ")
}
