package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApigDataSourceVo 创建据连接结构体信息
type ApigDataSourceVo struct {

	// 数据连接名称
	DwName string `json:"dw_name"`

	// 数据连接类型
	DwType string `json:"dw_type"`

	// 连接动态变化配置项，每种连接略有区别，建议在界面进行调试
	DwConfig *interface{} `json:"dw_config"`

	// 代理id（若使用代理连接则必填）
	AgentId *string `json:"agent_id,omitempty"`

	// 代理名称id（若使用代理连接则必填）
	AgentName *string `json:"agent_name,omitempty"`

	// 0：开发模式 1：生产模式。默认为0
	EnvType *int32 `json:"env_type,omitempty"`

	// 1：cdm 2：数据架构 4:数据开发 8：数据质量 16：数据目录 32：数据安全 64：数据服务
	SupportService *int32 `json:"supportService,omitempty"`
}

func (o ApigDataSourceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigDataSourceVo struct{}"
	}

	return strings.Join([]string{"ApigDataSourceVo", string(data)}, " ")
}
