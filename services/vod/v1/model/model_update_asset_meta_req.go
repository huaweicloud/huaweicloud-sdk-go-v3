package model

import (
	"encoding/json"

	"strings"
)

type UpdateAssetMetaReq struct {
	// 媒体ID<br/>

	AssetId string `json:"asset_id"`
	// 媒体标题<br/>

	Title *string `json:"title,omitempty"`
	// 视频描述<br/>

	Description *string `json:"description,omitempty"`
	// 媒资分类id<br/>

	CategoryId *int32 `json:"category_id,omitempty"`
	// 视频标签<br/>

	Tags *string `json:"tags,omitempty"`
	// 媒资所在文件夹id

	FolderName *string `json:"folder_name,omitempty"`
	// 自定义元数据<br/>

	CustomMetadata map[string]interface{} `json:"custom_metadata,omitempty"`
	// 权限<br/>

	Privilege *string `json:"privilege,omitempty"`
}

func (o UpdateAssetMetaReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAssetMetaReq struct{}"
	}

	return strings.Join([]string{"UpdateAssetMetaReq", string(data)}, " ")
}
