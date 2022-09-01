package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 发起实例诊断请求体
type CreateDiagnosisTaskBody struct {

	// 诊断开始时间。UNIX时间戳，单位毫秒。
	BeginTime string `json:"begin_time" xml:"begin_time"`

	// 诊断结束时间。UNIX时间戳，单位毫秒。
	EndTime string `json:"end_time" xml:"end_time"`

	// 诊断节点IP列表。默认诊断所有节点。 非读写分离实例查询方法如下：   - 方法一：参考[查看实例信息](https://support.huaweicloud.com/usermanual-dcs/dcs-ug-0312016.html)。   - 方法二：调用[查询指定实例](https://support.huaweicloud.com/api-dcs/ShowInstance.html)查询。  读写分离实例查询方法：调用[查询分片信息](https://support.huaweicloud.com/api-dcs/ListGroupReplicationInfo.html#ListGroupReplicationInfo__response_InstanceReplicationListInfo)。
	NodeIpList *[]string `json:"node_ip_list,omitempty" xml:"node_ip_list"`
}

func (o CreateDiagnosisTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDiagnosisTaskBody struct{}"
	}

	return strings.Join([]string{"CreateDiagnosisTaskBody", string(data)}, " ")
}
