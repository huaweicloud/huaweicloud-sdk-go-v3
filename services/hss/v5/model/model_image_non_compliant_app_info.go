package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageNonCompliantAppInfo struct {

	// **参数解释**: id **取值范围**: 字符长度0-128位
	Id *string `json:"id,omitempty"`

	// **参数解释**: 不合规软件名称 **取值范围**: 字符长度0-128位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 不合规软件路径 **取值范围**: 字符长度0-128位
	AppPath *string `json:"app_path,omitempty"`

	// **参数解释**: 不合规软件版本 **取值范围**: 字符长度0-64位
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**: 层digest **取值范围**: 字符长度0-128位
	LayerDigest *string `json:"layer_digest,omitempty"`
}

func (o ImageNonCompliantAppInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageNonCompliantAppInfo struct{}"
	}

	return strings.Join([]string{"ImageNonCompliantAppInfo", string(data)}, " ")
}
