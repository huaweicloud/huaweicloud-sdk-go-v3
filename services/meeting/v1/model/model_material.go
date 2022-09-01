package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 素材信息
type Material struct {

	// 素材ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 更新者
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty" xml:"lastUpdatedBy"`

	// 更新时间
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime"`

	// 素材名称
	MaterialName *string `json:"materialName,omitempty" xml:"materialName"`

	// 素材分辨率
	MaterialResolution *string `json:"materialResolution,omitempty" xml:"materialResolution"`

	// 素材大小（含单位）
	MaterialSizeStr *string `json:"materialSizeStr,omitempty" xml:"materialSizeStr"`

	// 素材云盘存储文件下载地址
	FilePath *string `json:"filePath,omitempty" xml:"filePath"`
}

func (o Material) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Material struct{}"
	}

	return strings.Join([]string{"Material", string(data)}, " ")
}
