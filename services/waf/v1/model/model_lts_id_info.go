package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 日志组与日志流id信息，对应云日志服务（lts）创建的日志组和日志流id。
type LtsIdInfo struct {

	// 日志组id
	LtsGroupId *string `json:"ltsGroupId,omitempty"`

	// 访问日志流id
	LtsAccessStreamID *string `json:"ltsAccessStreamID,omitempty"`

	// 攻击日志流id
	LtsAttackStreamID *interface{} `json:"ltsAttackStreamID,omitempty"`
}

func (o LtsIdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsIdInfo struct{}"
	}

	return strings.Join([]string{"LtsIdInfo", string(data)}, " ")
}
