package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAcceleratorOption 创建全球加速器实例的详细信息。
type CreateAcceleratorOption struct {

	// 全球加速器名称，取值范围：1~64个字符之间，只能由数字、字母、中划线和中文组成。
	Name string `json:"name"`

	// 全球加速器描述信息，取值范围：0~255个字符之间，禁止输入字符：<>。
	Description *string `json:"description,omitempty"`

	// 全球加速器IP列表。
	IpSets []CreateAcceleratorOptionIpSets `json:"ip_sets"`

	// 租户的企业项目ID，最大长度36个字符，带\"-\"连字符的UUID格式，或者是字符串\"0\"。\"0\"表示默认企业项目。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`
}

func (o CreateAcceleratorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAcceleratorOption struct{}"
	}

	return strings.Join([]string{"CreateAcceleratorOption", string(data)}, " ")
}
