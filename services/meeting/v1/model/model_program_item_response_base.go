package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节目素材信息。
type ProgramItemResponseBase struct {

	// 素材ID。
	MaterialId *string `json:"materialId,omitempty"`

	// 素材名称。
	MaterialName *string `json:"materialName,omitempty"`

	// 素材云盘文件下载路径。
	FilePath *string `json:"filePath,omitempty"`

	// 播放时长。
	PlayTime *int32 `json:"playTime,omitempty"`
}

func (o ProgramItemResponseBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgramItemResponseBase struct{}"
	}

	return strings.Join([]string{"ProgramItemResponseBase", string(data)}, " ")
}
