package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppWhitelistAbnormalInfo 学习异常原因信息
type AppWhitelistAbnormalInfo struct {

	// 异常类型
	AbnormalType *int32 `json:"abnormal_type,omitempty"`

	// 异常描述
	AbnormalDescription *string `json:"abnormal_description,omitempty"`
}

func (o AppWhitelistAbnormalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppWhitelistAbnormalInfo struct{}"
	}

	return strings.Join([]string{"AppWhitelistAbnormalInfo", string(data)}, " ")
}
