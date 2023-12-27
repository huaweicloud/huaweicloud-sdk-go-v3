package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReviewConfig 内容审核配置
type ReviewConfig struct {

	// 免审核。 目前仅白名单用户可使用此参数，非白名单用户跟随系统策略审核。
	NoNeedReview *bool `json:"no_need_review,omitempty"`
}

func (o ReviewConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReviewConfig struct{}"
	}

	return strings.Join([]string{"ReviewConfig", string(data)}, " ")
}
