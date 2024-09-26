package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BackgroundInfoVo struct {

	// 项目id
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// cover文件名称
	CoverFileName *string `json:"cover_file_name,omitempty"`

	// 背景文件名称
	BackgroundFileName *string `json:"background_file_name,omitempty"`

	// logo文件名称
	LogoFileName *string `json:"logo_file_name,omitempty"`
}

func (o BackgroundInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackgroundInfoVo struct{}"
	}

	return strings.Join([]string{"BackgroundInfoVo", string(data)}, " ")
}
