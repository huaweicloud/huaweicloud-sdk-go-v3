package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CopyImageCrossRegionRequestBody struct {

	// IMS服务委托名称。
	AgencyName string `json:"agency_name" xml:"agency_name"`

	// 镜像描述信息。支持字母、数字、中文等，不支持回车、<、 >，长度不能超过1024个字符。默认为空。
	Description *string `json:"description,omitempty" xml:"description"`

	// 镜像名称
	Name string `json:"name" xml:"name"`

	// 目的区域的项目名称。
	ProjectName string `json:"project_name" xml:"project_name"`

	// 目的区域的Region ID。
	Region string `json:"region" xml:"region"`
}

func (o CopyImageCrossRegionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageCrossRegionRequestBody struct{}"
	}

	return strings.Join([]string{"CopyImageCrossRegionRequestBody", string(data)}, " ")
}
