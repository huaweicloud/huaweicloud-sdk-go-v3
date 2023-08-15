package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VarInfo struct {

	// 动参变量占位符名称。  > 格式：#p_N#（N表示第几个参数，与模板文件占位名称保持一致），如：#p_1#。
	Name *string `json:"name,omitempty"`

	// 动参变量类型。  - 类型为文字：txt - 类型为图片：jpg/jpeg/png/gif - 类型为音频：mp3/wav - 类型为视频：mp4/3gp
	Type *string `json:"type,omitempty"`
}

func (o VarInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VarInfo struct{}"
	}

	return strings.Join([]string{"VarInfo", string(data)}, " ")
}
