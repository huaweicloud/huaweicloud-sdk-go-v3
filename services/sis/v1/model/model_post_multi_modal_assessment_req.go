package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type PostMultiModalAssessmentReq struct {
	Config *MultiModalConfig `json:"config"`

	// 视频数据，Base64编码，要求Base64编码后大小不超过10M。  注意评测接口使用次数定义为：每8秒的视频作为一次，不足8秒按一次计算。例如传入4秒或8秒的视频，都算作使用一次，传入9秒的视频则视为调用2次。
	VideoData string `json:"video_data"`

	// 被评估视频和语音数据对应的试题文本，长度不可超过256字节。
	RefText string `json:"ref_text"`
}

func (o PostMultiModalAssessmentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostMultiModalAssessmentReq struct{}"
	}

	return strings.Join([]string{"PostMultiModalAssessmentReq", string(data)}, " ")
}
