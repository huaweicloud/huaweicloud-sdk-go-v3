package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 扫描任务配置
type TaskSettingsTaskConfig struct {
	// 扫描模式:   * fast - 快速扫描   * normal - 标准扫描   * deep - 深度扫描

	ScanMode *TaskSettingsTaskConfigScanMode `json:"scan_mode,omitempty"`
	// 是否进行端口扫描

	PortScan *bool `json:"port_scan,omitempty"`
	// 是否进行弱密码扫描

	WeakPwdScan *bool `json:"weak_pwd_scan,omitempty"`
	// 是否进行CVE漏洞扫描

	CveCheck *bool `json:"cve_check,omitempty"`
	// 是否进行网站内容合规文字检测

	TextCheck *bool `json:"text_check,omitempty"`
	// 是否进行网站内容合规图片检测

	PictureCheck *bool `json:"picture_check,omitempty"`
	// 是否进行网站挂马检测

	MaliciousCode *bool `json:"malicious_code,omitempty"`
	// 是否进行链接健康检测（死链、暗链、恶意外链）

	MaliciousLink *bool `json:"malicious_link,omitempty"`
}

func (o TaskSettingsTaskConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskSettingsTaskConfig struct{}"
	}

	return strings.Join([]string{"TaskSettingsTaskConfig", string(data)}, " ")
}

type TaskSettingsTaskConfigScanMode struct {
	value string
}

type TaskSettingsTaskConfigScanModeEnum struct {
	FAST   TaskSettingsTaskConfigScanMode
	NORMAL TaskSettingsTaskConfigScanMode
	DEEP   TaskSettingsTaskConfigScanMode
}

func GetTaskSettingsTaskConfigScanModeEnum() TaskSettingsTaskConfigScanModeEnum {
	return TaskSettingsTaskConfigScanModeEnum{
		FAST: TaskSettingsTaskConfigScanMode{
			value: "fast",
		},
		NORMAL: TaskSettingsTaskConfigScanMode{
			value: "normal",
		},
		DEEP: TaskSettingsTaskConfigScanMode{
			value: "deep",
		},
	}
}

func (c TaskSettingsTaskConfigScanMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskSettingsTaskConfigScanMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
