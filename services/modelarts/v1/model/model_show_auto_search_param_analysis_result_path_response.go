package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchParamAnalysisResultPathResponse Response Object
type ShowAutoSearchParamAnalysisResultPathResponse struct {

	// 超参敏感度分析图像的保存路径。
	FilePath       *string `json:"file_path,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAutoSearchParamAnalysisResultPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchParamAnalysisResultPathResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchParamAnalysisResultPathResponse", string(data)}, " ")
}
