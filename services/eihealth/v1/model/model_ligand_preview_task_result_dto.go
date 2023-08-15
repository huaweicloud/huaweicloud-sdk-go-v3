package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LigandPreviewTaskResultDto 配体预览任务结果
type LigandPreviewTaskResultDto struct {

	// 已知的配体数量
	Count int32 `json:"count"`

	// 预览配体信息列表
	Ligands []LigandPreviewInfoDto `json:"ligands"`

	// 文件中是否还有更多配体没有处理；即当前数量是否表示整个文件的配体数量，若该值为false则表该配体文件含有count数量个配体；若该值为true则表示改配体文件含有大于count数量个配体（即count不完全统计）；例如：count=100且has_more=true，则前端可显示该配体文件含有\"100+个\"小分子
	HasMore bool `json:"has_more"`
}

func (o LigandPreviewTaskResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LigandPreviewTaskResultDto struct{}"
	}

	return strings.Join([]string{"LigandPreviewTaskResultDto", string(data)}, " ")
}
