package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 请求参数
type CopyImageInRegionRequestBody struct {
	// 加密密钥。默认为空。

	CmkId *string `json:"cmk_id,omitempty"`
	// 镜像描述信息。_description参数说明请参考镜像属性。支持字母、数字、中文等，不支持回车、<、 >，长度不能超过1024个字符。默认为空。

	Description *string `json:"description,omitempty"`
	// 表示当前镜像所属的企业项目。 取值为0或无该值，表示属于default企业项目。 取值为UUID，表示属于该UUID对应的企业项目。关于企业项目ID的获取及企业项目特性的详细信息，请参考《企业管理用户指南》。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 镜像名称

	Name string `json:"name"`
}

func (o CopyImageInRegionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageInRegionRequestBody struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionRequestBody", string(data)}, " ")
}
