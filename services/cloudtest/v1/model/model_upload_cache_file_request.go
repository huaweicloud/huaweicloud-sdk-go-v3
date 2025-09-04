package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadCacheFileRequest Request Object
type UploadCacheFileRequest struct {

	// 项目uuid
	ProjectId string `json:"project_id"`

	// 是否覆盖同名文件
	Override *bool `json:"override,omitempty"`

	// 附件挂载资源类型
	ParentType *string `json:"parent_type,omitempty"`

	// 附件挂载资源Uri
	ParentUri *string `json:"parent_uri,omitempty"`

	Body *UploadCacheFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadCacheFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadCacheFileRequest struct{}"
	}

	return strings.Join([]string{"UploadCacheFileRequest", string(data)}, " ")
}
