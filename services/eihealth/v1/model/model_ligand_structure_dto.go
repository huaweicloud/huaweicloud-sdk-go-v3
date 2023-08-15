package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LigandStructureDto 配体3D结构
type LigandStructureDto struct {

	// 配体格式，即文件后缀名
	Format string `json:"format"`

	// 是否压缩
	Compressed *bool `json:"compressed,omitempty"`

	// 结构数据，如压缩则需要解码、解压处理（ASCII Encode -> Base64 Decode -> GZip Inflate -> UTF-8 Decode）以得到原始字符串；如未压缩则为原始字符串
	Data string `json:"data"`
}

func (o LigandStructureDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LigandStructureDto struct{}"
	}

	return strings.Join([]string{"LigandStructureDto", string(data)}, " ")
}
