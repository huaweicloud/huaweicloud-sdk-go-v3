package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 导入应用参数
type AppSrcReq struct {

	// 目标应用名称 取值范围：长度为[1,56]，以大小写字母开头，允许出现中划线(-)、下划线(_)、小写字母和数字，且必须以大小写字母或数字结尾。
	DestinationAppName string `json:"destination_app_name"`

	// 目标应用版本 取值范围：长度[1,24]，以小写字母或数字或大写字母开头，允许出现中划线，必须以小写字母或数字或大写字母结尾。
	DestinationAppVersion string `json:"destination_app_version"`

	// 源应用id
	SourceAppId string `json:"source_app_id"`
}

func (o AppSrcReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppSrcReq struct{}"
	}

	return strings.Join([]string{"AppSrcReq", string(data)}, " ")
}
