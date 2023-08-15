package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListParseDomainDetailsRequest Request Object
type ListParseDomainDetailsRequest struct {

	// 域名
	DomainName string `json:"domain_name"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。默认情况下，fw_instance_Id为空时，返回帐号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`
}

func (o ListParseDomainDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListParseDomainDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListParseDomainDetailsRequest", string(data)}, " ")
}
