package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WiseEye struct {

	// 是否开启云眼告警配置
	Enable *string `json:"enable,omitempty"`

	// 云眼告警级别
	Level *string `json:"level,omitempty"`

	// 云眼告警区域，目前取值有：china（中国区），asiaAfricaLatin（亚非拉），europe（欧洲）
	RegionKey *string `json:"region_key,omitempty"`

	// 云眼告警id，对应云眼信息中的name
	ScopeId *string `json:"scope_id,omitempty"`

	// 云眼告警范围，对应云眼信息中的label
	ScopeName *string `json:"scope_name,omitempty"`
}

func (o WiseEye) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WiseEye struct{}"
	}

	return strings.Join([]string{"WiseEye", string(data)}, " ")
}
