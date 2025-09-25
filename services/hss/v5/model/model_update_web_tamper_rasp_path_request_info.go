package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateWebTamperRaspPathRequestInfo struct {

	// **参数解释**: 动态网页防篡改的Tomcat bin目录。 **约束限制**: 仅Linux服务器支持配置动态网页防篡改的Tomcat bin目录。 **取值范围**: 字符长度1-256位，必须以/开头，不能以/结尾，只能包含英文大小写字母，数字，下划线，中划线和点。 **默认取值**: 不涉及
	RaspPath string `json:"rasp_path"`
}

func (o UpdateWebTamperRaspPathRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWebTamperRaspPathRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateWebTamperRaspPathRequestInfo", string(data)}, " ")
}
