package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AdministrationAgencyNamePrimitiveTypeHolder struct {

	// 管理委托名称  资源编排服务使用该委托获取成员账号委托给管理账号的权限  当用户定义SELF_MANAGED权限类型时，administration_agency_name和administration_agency_urn 必须有且只有一个存在。  推荐用户在使用v5委托时给与administration_agency_urn，administration_agency_name只支持接收v3委托名称，若给与了v5委托名称，则会在部署模板时失败。  当用户使用SERVICE_MANAGED权限类型时，指定该参数将报错400。  [创建委托及授权方式](https://support.huaweicloud.com/usermanual-iam/iam_06_0002.html)
	AdministrationAgencyName *string `json:"administration_agency_name,omitempty"`
}

func (o AdministrationAgencyNamePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdministrationAgencyNamePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"AdministrationAgencyNamePrimitiveTypeHolder", string(data)}, " ")
}
