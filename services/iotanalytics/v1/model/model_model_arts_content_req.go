package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelArts数据源配置内容
type ModelArtsContentReq struct {

	// 服务名称
	ServiceName string `json:"service_name"`

	// 访问地址
	AccessAddress string `json:"access_address"`

	// 校验参数
	VerifyBody string `json:"verify_body"`

	// 租户的AK
	Ak string `json:"ak"`

	// 租户的SK
	Sk string `json:"sk"`

	// 项目id
	ProjectId string `json:"project_id"`
}

func (o ModelArtsContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelArtsContentReq struct{}"
	}

	return strings.Join([]string{"ModelArtsContentReq", string(data)}, " ")
}
