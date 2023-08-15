package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImmediateJobsRequest Request Object
type ListImmediateJobsRequest struct {

	// 语言。
	XLanguage *string `json:"X-Language,omitempty"`

	// 任务执行状态。 取值： - 值为“Running”，表示任务正在执行。 - 值为“Completed”，表示任务执行成功。 - 值为“Failed”，表示任务执行失败。 - 值为“Pending”，表示任务未执行。
	Status *string `json:"status,omitempty"`

	// 任务名称。取值有：  - \"CreateGaussDBforMySQLInstance\"表示创建实例。  - \"RestoreGaussDBforMySQLNewInstance\"表示恢复新实例。  - \"AddGaussDBforMySQLNodes\"表示添加节点。  - \"DeleteGaussDBforMySQLNode\"表示删除节点。  - \"RebootGaussDBforMySQLInstance\"表示重启实例。  - \"ModifyGaussDBforMySQLPort\"表示修改实例端口。  - \"ModifyGaussDBforMySQLSecurityGroup\"表示修改实例安全组。  - \"ResizeGaussDBforMySQLFlavor\"表示实例规格变更。  - \"SwitchoverGaussDBforMySQLMasterNode\"表示只读升主。  - \"GaussDBforMySQLBindEIP\"表示绑定弹性公网IP。  - \"GaussDBforMySQLUnbindEIP\"表示解绑弹性公网IP。  - \"RenameGaussDBforMySQLInstance\"表示修改实例名称。  - \"DeleteGaussDBforMySQLInstance\"表示删除实例集群。  - \"UpgradeGaussDBforMySQLDatabaseVersion\"表示版本升级。  - \"EnlargeGaussDBforMySQLProxy\"表示实例的数据库代理节点扩容。  - \"OpenGaussDBforMySQLProxy\"表示开启实例的数据库代理。  - \"CloseGaussDBforMySQLProxy\"表示关闭实例的数据库代理。  - \"GaussdbforMySQLModifyProxyIp\"表示修改数据库代理ip。  - \"ScaleGaussDBforMySQLProxy\"表示实例的数据库代理节点规格变更。  - \"GaussDBforMySQLModifyInstanceMetricExtend\"表示实例秒级监控。  - \"GaussDBforMySQLModifyInstanceDataVip\"表示修改实例数据Vip。  - \"GaussDBforMySQLSwitchSSL\"表示切换实例SSL开关。  - \"GaussDBforMySQLModifyProxyConsist\"表示修改代理一致性。  - \"GaussDBforMySQLModifyProxyWeight\"表示修改代理权重。
	JobName *string `json:"job_name,omitempty"`

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为1，必须为数字，不能为负数。
	Offset *string `json:"offset,omitempty"`

	// 查询记录数。默认为10，取值为10、20、50。
	Limit *string `json:"limit,omitempty"`

	// 起始时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListImmediateJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImmediateJobsRequest struct{}"
	}

	return strings.Join([]string{"ListImmediateJobsRequest", string(data)}, " ")
}
