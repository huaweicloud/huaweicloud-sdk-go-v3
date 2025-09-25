package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BlockedIpResponseInfo 已拦截IP详情
type BlockedIpResponseInfo struct {

	// **参数解释**： 主机ID **取值范围**： 字符长度1-64位
	HostId string `json:"host_id"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName string `json:"host_name"`

	// 攻击源IP
	SrcIp string `json:"src_ip"`

	// **参数解释**： 登录类型 **约束限制**: 不涉及 **取值范围**: - mysql：mysql服务。 - rdp：rdp服务。 - ssh：ssh服务。 - vsftp：vsftp服务。  **默认取值**: 不涉及
	LoginType string `json:"login_type"`

	// 拦截次数
	InterceptNum int32 `json:"intercept_num"`

	// **参数解释**： 拦截状态 **约束限制**: 不涉及 **取值范围**: - intercepted：已拦截。 - canceled：已解除拦截。 - cancelling：待解除拦截。  **默认取值**: 不涉及
	InterceptStatus BlockedIpResponseInfoInterceptStatus `json:"intercept_status"`

	// 开始拦截时间，毫秒
	BlockTime int64 `json:"block_time"`

	// 最近拦截时间，毫秒
	LatestTime int64 `json:"latest_time"`
}

func (o BlockedIpResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlockedIpResponseInfo struct{}"
	}

	return strings.Join([]string{"BlockedIpResponseInfo", string(data)}, " ")
}

type BlockedIpResponseInfoInterceptStatus struct {
	value string
}

type BlockedIpResponseInfoInterceptStatusEnum struct {
	INTERCEPTED BlockedIpResponseInfoInterceptStatus
	CANCELED    BlockedIpResponseInfoInterceptStatus
	CANCELLING  BlockedIpResponseInfoInterceptStatus
}

func GetBlockedIpResponseInfoInterceptStatusEnum() BlockedIpResponseInfoInterceptStatusEnum {
	return BlockedIpResponseInfoInterceptStatusEnum{
		INTERCEPTED: BlockedIpResponseInfoInterceptStatus{
			value: "intercepted",
		},
		CANCELED: BlockedIpResponseInfoInterceptStatus{
			value: "canceled",
		},
		CANCELLING: BlockedIpResponseInfoInterceptStatus{
			value: "cancelling",
		},
	}
}

func (c BlockedIpResponseInfoInterceptStatus) Value() string {
	return c.value
}

func (c BlockedIpResponseInfoInterceptStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BlockedIpResponseInfoInterceptStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
