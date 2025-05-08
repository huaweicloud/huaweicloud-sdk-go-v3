package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateObjectRetrievalRequestBody obs桶取回请求消息
type CreateObjectRetrievalRequestBody struct {
	ObsInfo *ObsInfo `json:"obs_info"`

	// 临时恢复时间，临时恢复会产生一个标准存储副本，副本的存在时间。单位：天，默认1天。
	Days *int32 `json:"days,omitempty"`

	// 回调地址，为空则不回调
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 自定义上下文，回调时会回调给用户，透传信息
	SessionContext *string `json:"session_context,omitempty"`
}

func (o CreateObjectRetrievalRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateObjectRetrievalRequestBody struct{}"
	}

	return strings.Join([]string{"CreateObjectRetrievalRequestBody", string(data)}, " ")
}
