package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigDataSourcesVo 创建据连接结构体列表信息
type ApigDataSourcesVo struct {

	// 企业模式空间下的数据连接还是简单模式空间下的连接,0:简单模式，1：企业模式
	Mode *int32 `json:"mode,omitempty"`

	// 连接是否可见,0：不可见，1：可见
	Visible *int32 `json:"visible,omitempty"`

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
