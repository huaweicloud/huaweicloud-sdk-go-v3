package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileType 文件类型，包含如下:   - 0 ：全部   - 1 : 可执行   - 2 : 压缩   - 3 : 脚本   - 4 : 文档   - 5 : 图片   - 6 : 音视频
type FileType struct {
}

func (o FileType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileType struct{}"
	}

	return strings.Join([]string{"FileType", string(data)}, " ")
}
