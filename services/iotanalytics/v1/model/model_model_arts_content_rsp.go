package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SMN数据源配置内容
type ModelArtsContentRsp struct {

	// 服务名称
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName"`

	// 访问地址
	AccessAddress *string `json:"accessAddress,omitempty" xml:"accessAddress"`

	// 校验参数
	VerifyBody *string `json:"verifyBody,omitempty" xml:"verifyBody"`

	// 租户的AK
	Ak *string `json:"ak,omitempty" xml:"ak"`

	// 租户的SK
	Sk *string `json:"sk,omitempty" xml:"sk"`

	// 项目id
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`
}

func (o ModelArtsContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelArtsContentRsp struct{}"
	}

	return strings.Join([]string{"ModelArtsContentRsp", string(data)}, " ")
}
