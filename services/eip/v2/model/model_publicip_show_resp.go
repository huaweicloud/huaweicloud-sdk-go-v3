/*
    * EIP
    *
    * 云服务接口
    *
*/

package model
import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
)

// 弹性公网IP列表返回体
type PublicipShowResp struct {
	// 弹性公网IP对应带宽ID
	BandwidthId string `json:"bandwidth_id,omitempty"`
	// 带宽名称
	BandwidthName string `json:"bandwidth_name,omitempty"`
	// 表示共享带宽或者独享带宽  取值范围：PER，WHOLE。  WHOLE表示共享带宽  PER表示独享带宽  约束：其中IPv6暂不支持WHOLE类型带宽。
	BandwidthShareType string `json:"bandwidth_share_type,omitempty"`
	// 带宽大小，单位为Mbit/s。
	BandwidthSize int32 `json:"bandwidth_size,omitempty"`
	// 弹性公网IP申请时间（UTC）
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
	// 企业项目ID。最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。  创建弹性公网IP时，给弹性公网IP绑定企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	// 弹性公网IP唯一标识
	Id string `json:"id,omitempty"`
	// 功能说明：端口id。  约束：只有绑定了的弹性公网IP查询才会返回该参数
	PortId string `json:"port_id,omitempty"`
	// 功能说明：绑定弹性公网IP的私有IP地址  约束：只有绑定了的弹性公网IP查询才会返回该参数
	PrivateIpAddress string `json:"private_ip_address,omitempty"`
	Profile *ProfileResp `json:"profile,omitempty"`
	// IPv4时是申请到的弹性公网IP地址，IPv6时是IPv6地址对应的IPv4地址
	PublicIpAddress string `json:"public_ip_address,omitempty"`
	// 功能说明：弹性公网IP的状态  取值范围：冻结FREEZED，绑定失败BIND_ERROR，绑定中BINDING，释放中PENDING_DELETE， 创建中PENDING_CREATE，创建中NOTIFYING，释放中NOTIFY_DELETE，更新中PENDING_UPDATE， 未绑定DOWN ，绑定ACTIVE，绑定ELB，绑定VPN，失败ERROR。
	Status string `json:"status,omitempty"`
	// 项目ID
	TenantId string `json:"tenant_id,omitempty"`
	// 弹性公网IP的类型
	Type string `json:"type,omitempty"`
	// IPv4时无此字段，IPv6时为申请到的弹性公网IP地址
	PublicIpv6Address string `json:"public_ipv6_address,omitempty"`
	// IP版本信息，取值范围是4和6  4：表示IPv4  6：表示IPv6
	IpVersion int32 `json:"ip_version,omitempty"`
}
