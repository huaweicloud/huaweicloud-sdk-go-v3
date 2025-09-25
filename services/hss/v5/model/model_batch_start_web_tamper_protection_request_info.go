package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartWebTamperProtectionRequestInfo 开启网页防篡改防护信息
type BatchStartWebTamperProtectionRequestInfo struct {

	// **参数解释**: 需要开启防护的服务器ID列表，仅支持填写未开启网页防篡改防护的服务器ID，已开启网页防篡改防护的服务器可使用UpdateWebTamperHostPolicy接口进行修改策略。 **约束限制** : 仅支持填写未开启网页防篡改防护的服务器ID，且Linux服务器和Windows服务器不可同时填写，需分批开启。 **取值范围**: 最少1条，最多20000条 **默认取值** : 不涉及
	HostIdList []string `json:"host_id_list"`

	// **参数解释**: 计费模式 **约束限制**: 不涉及 **取值范围**: - packet_cycle: 包年/包月，可填写resource_id。 - on_demand: 按需计费，无需填写resource_id。  **默认取值**: on_demand
	ChargingMode *string `json:"charging_mode,omitempty"`

	// **参数解释**: 资源ID，即网页防篡改配额的配额ID，当charging_mode选择packet_cycle时可填写该字段，表示使用一个指定配额，也可不填写该字段，表示随机选择符合的配额。 **约束限制** : 不涉及 **取值范围**: 字符长度0-64位 **默认取值** : 不涉及
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**： 资源标签列表，仅计费模式选择按需计费时支持填写。 **约束限制**: 仅计费模式选择按需计费时支持填写。 **取值范围**: 最少0条，最多2097152条 **默认取值**: 不涉及
	Tags *[]TagInfo `json:"tags,omitempty"`

	ProtectDirInfo *WebTamperProtectDirRequestInfo `json:"protect_dir_info"`

	// **参数解释**: 定时开关设置状态 **约束限制**: 不涉及 **取值范围**: - True ：开启定时关闭防护功能，必须填写timing_off_config_info。 - False ：关闭定时关闭防护功能，无需填写timing_off_config_info。  **默认取值**: False
	EnableTimingOff *bool `json:"enable_timing_off,omitempty"`

	TimingOffConfigInfo *WebTamperTimingOffConfigInfoRequestInfo `json:"timing_off_config_info,omitempty"`

	// **参数解释**: 动态网页防篡改开启状态，仅Linux服务器支持。 **约束限制**: 仅Linux服务器支持开启动态网页防篡改，Windows服务器不可填写该字段。 **取值范围**: - True ：开启动态网页防篡改，必须填写rasp_path。 - False ：关闭动态网页防篡改，无需填写rasp_path。  **默认取值**: False
	EnableRaspProtect *bool `json:"enable_rasp_protect,omitempty"`

	// **参数解释**: 动态网页防篡改的Tomcat bin目录，仅Linux服务器支持。 **约束限制**: 仅Linux服务器支持配置动态网页防篡改的Tomcat bin目录，Windows服务器不可填写该字段。 **取值范围**: 字符长度1-256位，必须以/开头，不能以/结尾，只能包含英文大小写字母，数字，下划线，中划线和点。 **默认取值**: 不涉及
	RaspPath *string `json:"rasp_path,omitempty"`

	// **参数解释**: 特权进程开启状态 **约束限制**: 不涉及 **取值范围**: - True ：开启特权进程，必须填写privileged_process_info。 - False ：关闭特权进程，无需填写privileged_process_info。  **默认取值**: False
	EnablePrivilegedProcess *bool `json:"enable_privileged_process,omitempty"`

	PrivilegedProcessInfo *WebTamperPrivilegedProcessRequestInfo `json:"privileged_process_info,omitempty"`
}

func (o BatchStartWebTamperProtectionRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartWebTamperProtectionRequestInfo struct{}"
	}

	return strings.Join([]string{"BatchStartWebTamperProtectionRequestInfo", string(data)}, " ")
}
