package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GlanceUpdateImageMemberRequestBody 镜像成员的状态。
type GlanceUpdateImageMemberRequestBody struct {

	// 镜像成员的状态。 取值如下： accepted：表示接受共享镜像。接受后，该镜像在用户镜像列表中可见，用户可以使用该镜像创建云服务器。 rejected：表示拒绝共享镜像。拒绝后，该镜像在用户镜像列表中不可见，但是，用户仍然可以使用该镜像创建云服务器。
	Status GlanceUpdateImageMemberRequestBodyStatus `json:"status"`

	// 存储库ID。 如果是CBR创建的整机镜像，则在接受该共享镜像时，为必选参数，需传入该值。 存储库ID可以从云备份服务控制台获取，或者参考《云备份接口参考》的“查询存储库列表”章节查询。
	VaultId *string `json:"vault_id,omitempty"`
}

func (o GlanceUpdateImageMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageMemberRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageMemberRequestBody", string(data)}, " ")
}

type GlanceUpdateImageMemberRequestBodyStatus struct {
	value string
}

type GlanceUpdateImageMemberRequestBodyStatusEnum struct {
	ACCEPTED GlanceUpdateImageMemberRequestBodyStatus
	REJECTED GlanceUpdateImageMemberRequestBodyStatus
}

func GetGlanceUpdateImageMemberRequestBodyStatusEnum() GlanceUpdateImageMemberRequestBodyStatusEnum {
	return GlanceUpdateImageMemberRequestBodyStatusEnum{
		ACCEPTED: GlanceUpdateImageMemberRequestBodyStatus{
			value: "accepted",
		},
		REJECTED: GlanceUpdateImageMemberRequestBodyStatus{
			value: "rejected",
		},
	}
}

func (c GlanceUpdateImageMemberRequestBodyStatus) Value() string {
	return c.value
}

func (c GlanceUpdateImageMemberRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageMemberRequestBodyStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
