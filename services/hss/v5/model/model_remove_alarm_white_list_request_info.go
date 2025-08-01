package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveAlarmWhiteListRequestInfo 删除告警白名单列表
type RemoveAlarmWhiteListRequestInfo struct {

	// 删除告警白名单详情
	DataList *[]AlarmWhiteListRequestInfo `json:"data_list,omitempty"`

	// **参数解释**: 是否需要恢复相关告警 **约束限制**: 不涉及 **取值范围**: - true ：恢复告警 - false ：不恢复告警 **默认取值**: false
	RestoreAlarm *bool `json:"restore_alarm,omitempty"`

	// **参数解释**: 是否删除所有白名单内容 **约束限制**: 不涉及 **取值范围**: - true ：删除所有白名单内容 - false ：不删除所有白名单内容 **默认取值**: false
	DeleteAll *bool `json:"delete_all,omitempty"`

	// 事件类型，包含如下:   - 1001 : 通用恶意软件   - 1002 : 病毒   - 1003 : 蠕虫   - 1004 : 木马   - 1005 : 僵尸网络   - 1006 : 后门   - 1010 : Rootkit   - 1011 : 勒索软件   - 1012 ：黑客工具   - 1015 : Webshell   - 1016 : 挖矿   - 1017 : 反弹Shell   - 2001 : 一般漏洞利用   - 2012 : 远程代码执行   - 2047 : Redis漏洞利用   - 2048 : Hadoop漏洞利用   - 2049 : MySQL漏洞利用   - 3002 : 文件提权   - 3003 : 进程提权   - 3004 : 关键文件变更   - 3005 : 文件/目录变更   - 3007 : 进程异常行为   - 3015 : 高危命令执行   - 3018 : 异常Shell   - 3027 : Crontab可疑任务   - 3029 ：系统安全防护被禁用   - 3030 ：备份删除   - 3031 ：异常注册表操作   - 3036 : 容器镜像阻断   - 4002 : 暴力破解   - 4004 : 异常登录   - 4006 : 非法系统账号   - 4014 : 用户账号添加   - 4020 : 用户密码窃取   - 6002 : 端口扫描   - 6003 : 主机扫描   - 13001 : Kubernetes事件删除   - 13002 : Pod异常行为   - 13003 : 枚举用户信息   - 13004 : 绑定集群用户角色
	EventType *int32 `json:"event_type,omitempty"`
}

func (o RemoveAlarmWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveAlarmWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"RemoveAlarmWhiteListRequestInfo", string(data)}, " ")
}
