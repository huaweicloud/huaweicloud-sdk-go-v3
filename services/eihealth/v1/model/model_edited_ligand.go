package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditedLigand struct {

	// 文件来源，仅支持RAW
	Source *string `json:"source,omitempty"`

	// 文件格式，仅支持CIF
	Format *string `json:"format,omitempty"`

	// 文件原始数据，仅数据源为RAW时提供
	Data string `json:"data"`
}

func (o EditedLigand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditedLigand struct{}"
	}

	return strings.Join([]string{"EditedLigand", string(data)}, " ")
}
