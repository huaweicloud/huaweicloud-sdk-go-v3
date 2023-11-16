package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OutPutInfoResult 结果
type OutPutInfoResult struct {
	PackageInfo *OutPutResult `json:"package_info,omitempty"`

	// 二方包信息
	PackageInfos map[string]OutPutResult `json:"package_infos,omitempty"`

	// 镜像包信息
	ImageInfos map[string]OutPutResult `json:"image_infos,omitempty"`
}

func (o OutPutInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OutPutInfoResult struct{}"
	}

	return strings.Join([]string{"OutPutInfoResult", string(data)}, " ")
}
