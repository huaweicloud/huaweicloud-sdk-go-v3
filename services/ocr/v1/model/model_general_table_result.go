package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type GeneralTableResult struct {
	// 文字区域数目。

	WordsRegionCount int32 `json:"words_region_count"`
	// 文字区域识别结果列表，输出顺序从左到右，先上后下。

	WordsRegionList []WordsRegionList `json:"words_region_list"`
	// 表格图像转换为excel的base64编码，图像中的文字和表格按位置写入excel。对返回的excel编码可用base64.b64decode解码并保存为.xlsx文件。

	Excel *string `json:"excel,omitempty"`
}

func (o GeneralTableResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeneralTableResult struct{}"
	}

	return strings.Join([]string{"GeneralTableResult", string(data)}, " ")
}
