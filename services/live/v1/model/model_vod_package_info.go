package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VodPackageInfo VOD 打包信息，包括DRM resourceID和转封装模板ID，模板ID通过VOD查询
type VodPackageInfo struct {

	// VOD 打包信息，转封装模板ID，模板ID通过VOD查询
	PackagingGroupId *string `json:"packaging_group_id,omitempty"`

	// DRM resourceID
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o VodPackageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VodPackageInfo struct{}"
	}

	return strings.Join([]string{"VodPackageInfo", string(data)}, " ")
}
