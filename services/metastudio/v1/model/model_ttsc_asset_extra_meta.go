package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TtscAssetExtraMeta 资产元数据。根据资产类型选择其中一个填写。
type TtscAssetExtraMeta struct {
	VoiceModelMeta *TtscVoiceModelAssetMeta `json:"voice_model_meta,omitempty"`
}

func (o TtscAssetExtraMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtscAssetExtraMeta struct{}"
	}

	return strings.Join([]string{"TtscAssetExtraMeta", string(data)}, " ")
}
