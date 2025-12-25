package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAgentFileRequest Request Object
type UploadAgentFileRequest struct {

	// **参数解释**： 空间ID，当前资源所属工作空间  **取值范围**： 由英文，数字，“-”，“_”组成，不超过64位字符
	WorkspaceId string `json:"workspace_id"`

	// 访问授权过期时间（天）
	Expires *int32 `json:"expires,omitempty"`

	// 是否是图片上传
	IsImage *bool `json:"is_image,omitempty"`

	Body *UploadAgentFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadAgentFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAgentFileRequest struct{}"
	}

	return strings.Join([]string{"UploadAgentFileRequest", string(data)}, " ")
}
