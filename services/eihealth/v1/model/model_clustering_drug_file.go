package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ClusteringDrugFile struct {
	Source *ClusteringFileSource `json:"source"`

	// 文件URL，当数据源为用户私有数据中心为项目路径，为公共数据场景时为obs地址。
	Url *string `json:"url,omitempty"`

	// 文件格式，支持SMI，仅数据源为RAW时提供。
	Format *string `json:"format,omitempty"`

	// 文件原始数据，仅数据源为RAW时提供。
	Data *string `json:"data,omitempty"`
}

func (o ClusteringDrugFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusteringDrugFile struct{}"
	}

	return strings.Join([]string{"ClusteringDrugFile", string(data)}, " ")
}
