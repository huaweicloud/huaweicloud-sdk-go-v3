package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FtMetric 训练指标结构体
type FtMetric struct {

	// 指标中文名称，如 训练指标、准确率，前端用作图例或列名
	NameCn string `json:"name_cn"`

	// 指标英文名称，如 train_loss、val_accuracy，前端用作图例或列名
	NameEn string `json:"name_en"`

	// 指标中文解释，如 训练指标，前端用作针对指标进行释义
	DesEn string `json:"des_en"`

	// 指标英文解释，如 train loss，前端用作针对指标进行释义
	DesCn string `json:"des_cn"`

	// 指标绘图类型，可选 line（折线图）或 pie（饼图）、tabel（表格）、scalar（单值），可扩展 image 等
	Type *string `json:"type,omitempty"`

	// 逻辑分组，如 training、validation、test，可扩展，用于前端分栏或过滤
	Group *string `json:"group,omitempty"`

	// 指定哪些数据点字段用于分组生成多个系列（如 [\"layer\",\"feature\"]）
	GroupBy *[]string `json:"group_by,omitempty"`

	// 明确指定用作 X 轴的数据点字段名（如 \"step\"、\"epoch\"、\"timestamp\"）
	XAxis *string `json:"x_axis,omitempty"`

	// 逻辑分组，如 表面loss，用于前端分组或过滤
	Tags *[]string `json:"tags,omitempty"`

	// 单位，如 %、samples/sec，仅用于展示
	Unit *string `json:"unit,omitempty"`

	// 数据点数组，严格按时间/步序升序排列
	Data *interface{} `json:"data"`
}

func (o FtMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FtMetric struct{}"
	}

	return strings.Join([]string{"FtMetric", string(data)}, " ")
}
