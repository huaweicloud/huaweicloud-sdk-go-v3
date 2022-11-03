package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VulnRule struct {

	// 问题文件
	FilePath *string `json:"file_path,omitempty"`

	// 特征信息
	IdentityInfo *string `json:"identity_info,omitempty"`
}

func (o VulnRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulnRule struct{}"
	}

	return strings.Join([]string{"VulnRule", string(data)}, " ")
}
