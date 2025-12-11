package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunImageSynchronizeResponse Response Object
type RunImageSynchronizeResponse struct {

	// **参数解释**: 错误编码 **取值范围**: 0（成功）、非0（失败）
	ErrorCode *int32 `json:"error_code,omitempty"`

	// **参数解释**: 错误描述 **取值范围**: 字符长度0-512位
	ErrorDescription *string `json:"error_description,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o RunImageSynchronizeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunImageSynchronizeResponse struct{}"
	}

	return strings.Join([]string{"RunImageSynchronizeResponse", string(data)}, " ")
}
