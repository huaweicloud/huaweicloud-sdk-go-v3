package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopUa top_ua配置
type TopUa struct {

	// 配置开关
	Enable *bool `json:"enable,omitempty"`
}

func (o TopUa) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUa struct{}"
	}

	return strings.Join([]string{"TopUa", string(data)}, " ")
}
