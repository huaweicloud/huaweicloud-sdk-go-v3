package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FinancialStatementResult
type FinancialStatementResult struct {

	// 识别出来的表格、文本区域个数。
	WordsRegionCount int32 `json:"words_region_count"`

	// 返回的表格、文本区域列表。输出顺序从左到右，从上到下。
	WordsRegionList []FinancialStatementWordsRegionList `json:"words_region_list"`

	// 表格图像转换为excel的base64编码，图像中的文字和表格按位置写入excel，可编辑。对返回的excel编码，可用base64.b64decode解码并保存为xlsx文件。
	Excel *string `json:"excel,omitempty"`

	ImageSize *FinancialStatementResultImageSize `json:"image_size,omitempty"`

	// 返回透视变换矩阵
	RectificationMatrix *[][]float32 `json:"rectification_matrix,omitempty"`
}

func (o FinancialStatementResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FinancialStatementResult struct{}"
	}

	return strings.Join([]string{"FinancialStatementResult", string(data)}, " ")
}
