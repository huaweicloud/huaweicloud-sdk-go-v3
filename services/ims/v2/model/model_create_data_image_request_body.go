package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 创建镜像请求体
type CreateDataImageRequestBody struct {
	// 创建加密镜像的用户主密钥，具体取值请参考《密钥管理服务用户指南》获取。

	CmkId *string `json:"cmk_id,omitempty"`
	// 镜像描述信息。_description参数说明请参考镜像属性。支持字母、数字、中文等，不支持回车、<、 >，长度不能超过1024个字符。默认为空。

	Description *string `json:"description,omitempty"`
	// 表示当前镜像所属的企业项目。取值为0或无该值，表示属于default企业项目；取值为UUID，表示属于该UUID对应的企业项目。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 新规范的镜像标签列表。默认为空。 tags和image_tags只能使用一个。

	ImageTags *[]ImageTag `json:"image_tags,omitempty"`
	// OBS桶中外部镜像文件地址。格式为<OBS桶名>:<OBS镜像文件名称>。 此处的OBS桶和镜像文件的存储类别必须是OBS标准存储。

	ImageUrl string `json:"image_url"`
	// 最小数据盘大小。取值范围40-2048GB。

	MinDisk int32 `json:"min_disk"`
	// 镜像名称。

	Name string `json:"name"`
	// 操作系统类型。只能是Windows、Linux二者之一，默认Linux。

	OsType *CreateDataImageRequestBodyOsType `json:"os_type,omitempty"`
	// 镜像标签列表。默认为空。 tags和image_tags只能使用一个。

	Tags *[]string `json:"tags,omitempty"`
}

func (o CreateDataImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImageRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDataImageRequestBody", string(data)}, " ")
}

type CreateDataImageRequestBodyOsType struct {
	value string
}

type CreateDataImageRequestBodyOsTypeEnum struct {
	WINDOWS CreateDataImageRequestBodyOsType
	LINUX   CreateDataImageRequestBodyOsType
}

func GetCreateDataImageRequestBodyOsTypeEnum() CreateDataImageRequestBodyOsTypeEnum {
	return CreateDataImageRequestBodyOsTypeEnum{
		WINDOWS: CreateDataImageRequestBodyOsType{
			value: "Windows",
		},
		LINUX: CreateDataImageRequestBodyOsType{
			value: "Linux",
		},
	}
}

func (c CreateDataImageRequestBodyOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDataImageRequestBodyOsType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
