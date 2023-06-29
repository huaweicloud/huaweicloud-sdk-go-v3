package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigDataSourcesVo 创建据连接结构体列表信息
type ApigDataSourcesVo struct {

	// 数据源结构体
	DataSourceVos *[]ApigDataSourceVo `json:"data_source_vos,omitempty"`
}

func (o ApigDataSourcesVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigDataSourcesVo struct{}"
	}

	return strings.Join([]string{"ApigDataSourcesVo", string(data)}, " ")
}
