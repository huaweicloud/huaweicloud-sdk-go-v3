package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateWholeImageRequestBody struct {
	// 镜像描述信息。 支持字母、数字、中文等，不支持回车、<、 >，长度不能超过1024个字符。

	Description *string `json:"description,omitempty"`
	// 表示当前镜像所属的企业项目。取值为0或无该值，表示属于default企业项目。取值为UUID，表示属于该UUID对应的企业项目。关于企业项目ID的获取及企业项目特性的详细信息，请参考《企业管理用户指南》。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 新规范的镜像标签列表。默认为空。tags和image_tags只能使用一个。

	ImageTags *[]TagKeyValue `json:"image_tags,omitempty"`
	// 弹性云服务器ID。使用弹性云服务器制作整机镜像时使用此参数且必填。 如果使用备份创建整机镜像，该参数应换成backup_id。 非必填的原因是需要兼容“使用备份创建整机镜像”和“使用弹性云服务器制作整机镜像”两种场景的body体。

	InstanceId *string `json:"instance_id,omitempty"`
	// 镜像名称。 名称的首尾字母不能为空格。 名称的长度至为1～128位。 名称包含以下4种字符： 大写字母 小写字母 数字 特殊字符包含-、.、_、空格和中文。

	Name string `json:"name"`
	// 镜像标签列表。tags和image_tags只能使用一个。

	Tags *[]string `json:"tags,omitempty"`
	// 使用云服务器备份创建整机镜像使用此参数且必填。 如果使用ECS创建整机镜像，则该参数应传为instance_id。 非必填的原因是需要兼容“使用备份创建整机镜像”和“使用弹性云服务器制作整机镜像”两种场景的body体。

	BackupId *string `json:"backup_id,omitempty"`
	// 使用备份创建整机镜像时，该字段区分是CBR服务的备份还是CSBS服务的备份，取值为：CBR/CSBS。 使用ECS创建整机镜像时，该字段不填

	WholeImageType *string `json:"whole_image_type,omitempty"`
	// 表示镜像支持的最大内存，单位为MB，默认不设置。

	MaxRam *int32 `json:"max_ram,omitempty"`
	// 表示镜像支持的最小内存，单位为MB，默认为0。

	MinRam *int32 `json:"min_ram,omitempty"`
	// 表示云服务器待加入的或已加入的存储库的ID。 使用云服务器创建整机镜像的过程为：先创建一个备份，再将备份创建为整机镜像。如果这个备份为CBR，vault_id为必填项；如果备份为CSBS，vault_id参数可不填。

	VaultId *string `json:"vault_id,omitempty"`
}

func (o CreateWholeImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWholeImageRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWholeImageRequestBody", string(data)}, " ")
}
