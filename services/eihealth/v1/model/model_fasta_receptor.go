package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FastaReceptor 受体文件（Fasta格式）
type FastaReceptor struct {
	Source *DrugFileSource `json:"source"`

	// 文件URL，当数据源为外部网络数据时为https地址；用户私有数据中心为项目路径、公共数据场景为obs地址
	Url *string `json:"url,omitempty"`

	// 文件格式，仅支持FASTA，仅数据源为RAW时提供
	Format *string `json:"format,omitempty"`

	// 文件原始数据，仅数据源为RAW时提供
	Data *string `json:"data,omitempty"`
}

func (o FastaReceptor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FastaReceptor struct{}"
	}

	return strings.Join([]string{"FastaReceptor", string(data)}, " ")
}
