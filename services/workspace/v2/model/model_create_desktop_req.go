package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建桌面请求。
type CreateDesktopReq struct {

	// 云桌面类型。  - DEDICATED：专属桌面。
	DesktopType CreateDesktopReqDesktopType `json:"desktop_type"`

	// 可用分区。将桌面创建到指定的可用分区。如果不指定则使用系统随机的可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 套餐ID。
	ProductId string `json:"product_id"`

	// 镜像类型。默认值为private。  - private：私有镜像。 - gold：公共镜像。
	ImageType string `json:"image_type"`

	// 镜像ID，用于私有镜像创建桌面场景，配合product_id使用。
	ImageId string `json:"image_id"`

	RootVolume *Volume `json:"root_volume"`

	// 数据盘列表。
	DataVolumes *[]Volume `json:"data_volumes,omitempty"`

	// 桌面对应的网卡信息，如果不指定则使用默认网卡。
	Nics *[]Nic `json:"nics,omitempty"`

	// 桌面使用的安全组，如果不指定则默认使用桌面代理中指定的安全组。
	SecurityGroups *[]SecurityGroupInfo `json:"security_groups,omitempty"`

	// 创建桌面使用的参数列表。长度为1-50。  当前不支持一批桌面不同配置，所有桌面的配置和第一台的一致，如果第一台未设置参数，则取外层的同名参数。
	Desktops []Desktop `json:"desktops"`

	// 创建成功后是否发送邮件通知桌面用户，默认为true。此参数仅在开通云桌面服务的domain_type为LOCAL_AD时有效，为LITE_AS时无效，因为LITE_AS首次创建桌面时必须发送邮件通知桌面用户修改登录密码。
	EmailNotification *bool `json:"email_notification,omitempty"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o CreateDesktopReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopReq struct{}"
	}

	return strings.Join([]string{"CreateDesktopReq", string(data)}, " ")
}

type CreateDesktopReqDesktopType struct {
	value string
}

type CreateDesktopReqDesktopTypeEnum struct {
	DEDICATED CreateDesktopReqDesktopType
}

func GetCreateDesktopReqDesktopTypeEnum() CreateDesktopReqDesktopTypeEnum {
	return CreateDesktopReqDesktopTypeEnum{
		DEDICATED: CreateDesktopReqDesktopType{
			value: "DEDICATED",
		},
	}
}

func (c CreateDesktopReqDesktopType) Value() string {
	return c.value
}

func (c CreateDesktopReqDesktopType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDesktopReqDesktopType) UnmarshalJSON(b []byte) error {
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
