package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CpuOptions 自定义CPU选项。
type CpuOptions struct {

	// CPU超线程数， 决定CPU是否开启超线程。取值范围：1，2。  - 1: 关闭超线程。 - 2: 打开超线程。  取值不能大于flavor中的   hw:cpu_threads，并且目标flavor配置需满足 \"hw:cpu_policy\": \"dedicated\",\"hw:cpu_threads\": \"2\"需要同时满足如下条件，才能设置为“关闭超线程”  - 只能在实例创建或者resize时指定。 - 只有目标flavor的extra_specs参数： - 存在“hw:cpu_policy”并取值为“dedicated”。 - 存在“hw:cpu_threads”并取值为“2”。
	HwcpuThreads *CpuOptionsHwcpuThreads `json:"hw:cpu_threads,omitempty"`
}

func (o CpuOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CpuOptions struct{}"
	}

	return strings.Join([]string{"CpuOptions", string(data)}, " ")
}

type CpuOptionsHwcpuThreads struct {
	value int32
}

type CpuOptionsHwcpuThreadsEnum struct {
	E_1 CpuOptionsHwcpuThreads
	E_2 CpuOptionsHwcpuThreads
}

func GetCpuOptionsHwcpuThreadsEnum() CpuOptionsHwcpuThreadsEnum {
	return CpuOptionsHwcpuThreadsEnum{
		E_1: CpuOptionsHwcpuThreads{
			value: 1,
		}, E_2: CpuOptionsHwcpuThreads{
			value: 2,
		},
	}
}

func (c CpuOptionsHwcpuThreads) Value() int32 {
	return c.value
}

func (c CpuOptionsHwcpuThreads) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CpuOptionsHwcpuThreads) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
