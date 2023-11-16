package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryJobNoticeItems struct {

	// 通知类型
	NoticeType *string `json:"notice_type,omitempty"`

	// 通知品种开启详情的表
	EnableMap map[string]bool `json:"enable_map,omitempty"`

	// 参数配置
	ParamConfig *string `json:"param_config,omitempty"`
}

func (o QueryJobNoticeItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryJobNoticeItems struct{}"
	}

	return strings.Join([]string{"QueryJobNoticeItems", string(data)}, " ")
}
