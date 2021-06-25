package model

import (
	"encoding/json"

	"strings"
)

type AssetProcessReq struct {
	// 媒体ID<br/>

	AssetId string `json:"asset_id"`
	// 转码模板组名称<br/>

	TemplateGroupName *string `json:"template_group_name,omitempty"`
	// 是否自动加密，取值[0，1] 0表示不加密，1表示需要加密。默认值为0 加密时必须要转码，且转码的输出是HLS。

	AutoEncrypt *int32 `json:"auto_encrypt,omitempty"`

	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`
	// 字幕文件id<br/>

	SubtitleId *[]int32 `json:"subtitle_id,omitempty"`
}

func (o AssetProcessReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssetProcessReq struct{}"
	}

	return strings.Join([]string{"AssetProcessReq", string(data)}, " ")
}
