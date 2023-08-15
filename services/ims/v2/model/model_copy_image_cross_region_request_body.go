package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CopyImageCrossRegionRequestBody struct {

	// IMS服务委托名称。
	AgencyName string `json:"agency_name"`

	// 镜像描述信息。支持字母、数字、中文等，不支持回车、<、 >，长度不能超过1024个字符。默认为空。
	Description *string `json:"description,omitempty"`

	// 镜像名称
	Name string `json:"name"`

	// 目的区域的项目名称。
	ProjectName string `json:"project_name"`

	// 目的区域的Region ID。
	Region string `json:"region"`

	// 存储库ID。如果是整机镜像，则在跨Region复制镜像时，为必选参数，需传入该值。
	VaultId *string `json:"vault_id,omitempty"`
}

func (o CopyImageCrossRegionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageCrossRegionRequestBody struct{}"
	}

	return strings.Join([]string{"CopyImageCrossRegionRequestBody", string(data)}, " ")
}
