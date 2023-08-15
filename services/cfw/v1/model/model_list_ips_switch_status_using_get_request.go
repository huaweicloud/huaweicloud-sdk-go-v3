package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpsSwitchStatusUsingGetRequest Request Object
type ListIpsSwitchStatusUsingGetRequest struct {

	// 防护对象id，是创建云防火墙后用于区分互联网边界防护和VPC边界防护的标志id，可通过调用查询防火墙实例接口获得，注意type为0的为互联网边界防护对象id，type为1的为VPC边界防护对象id。具体可参考APIExlorer和帮助中心FAQ。
	ObjectId string `json:"object_id"`

	// 企业项目id，用户支持企业项目后，由企业项目生成的id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防火墙实例id，创建云防火墙后用于标志防火墙由系统自动生成的标志id，可通过调用查询防火墙实例接口获得。具体可参考APIExlorer和帮助中心FAQ。默认情况下，fw_instance_Id为空时，返回帐号下第一个墙的信息；fw_instance_Id非空时，返回与fw_instance_Id对应墙的信息。
	FwInstanceId *string `json:"fw_instance_id,omitempty"`
}

func (o ListIpsSwitchStatusUsingGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpsSwitchStatusUsingGetRequest struct{}"
	}

	return strings.Join([]string{"ListIpsSwitchStatusUsingGetRequest", string(data)}, " ")
}
