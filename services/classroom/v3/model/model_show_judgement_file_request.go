package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJudgementFileRequest struct {
	// 文件id，可以从接口[ShowJudgementDetail](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=Classroom&api=ShowJudgementDetail)的响应中获取

	FileId string `json:"file_id"`
}

func (o ShowJudgementFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJudgementFileRequest struct{}"
	}

	return strings.Join([]string{"ShowJudgementFileRequest", string(data)}, " ")
}
