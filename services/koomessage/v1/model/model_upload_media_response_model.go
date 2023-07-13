package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadMediaResponseModel 上传素材返回信息。
type UploadMediaResponseModel struct {

	// 资源类型。 - 1：图片
	ResourceType *int32 `json:"resource_type,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源路径。
	ResourceUrl *string `json:"resource_url,omitempty"`
}

func (o UploadMediaResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadMediaResponseModel struct{}"
	}

	return strings.Join([]string{"UploadMediaResponseModel", string(data)}, " ")
}
