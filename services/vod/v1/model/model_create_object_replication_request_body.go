package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectReplicationRequestBody obs桶复制请求消息
type CreateObjectReplicationRequestBody struct {
	SrcInfo *ObsInfo `json:"src_info"`

	// 源文件类型
	FileType *string `json:"file_type,omitempty"`

	DestInfo *ObsInfo `json:"dest_info"`

	// 媒资分类id
	CategoryId *interface{} `json:"category_id,omitempty"`

	// 回调地址，为空则不回调
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 自定义上下文，回调时会回调给用户，透传信息
	SessionContext *string `json:"session_context,omitempty"`
}

func (o CreateObjectReplicationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectReplicationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateObjectReplicationRequestBody", string(data)}, " ")
}
