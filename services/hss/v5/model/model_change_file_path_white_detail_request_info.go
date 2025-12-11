package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeFilePathWhiteDetailRequestInfo 镜像的敏感信息文件路径白名单
type ChangeFilePathWhiteDetailRequestInfo struct {

	// **参数解释**: 自定义路径，选填，可编辑 **取值范围**: 不涉及
	CustomizedPath *[]string `json:"customized_path,omitempty"`
}

func (o ChangeFilePathWhiteDetailRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFilePathWhiteDetailRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeFilePathWhiteDetailRequestInfo", string(data)}, " ")
}
