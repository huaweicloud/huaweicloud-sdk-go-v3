package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryResourceResultDto struct {

	// 唯一标识若携带则以携带为准，企业内保证唯一，否则后台自动生成UUID。
	Id *string `json:"id,omitempty"`

	// 资源类型。 - VMR        - 云会议室 - CONF_CALL  - 会议并发数 - HARD_1080P - 1080P硬终端 - HARD_720P  - 720P硬终端 - SOFT       - 软终端用户数 - ROOM       - 大屏软终端 - LIVE       - 直播推流 - RECORD     - 录播空间 - HARD_THIRD_PARTY - 第三方硬终端帐号 - HUAWEI_VISION -智慧屏 - IDEA_HUB   - ideahub
	Type *string `json:"type,omitempty"`

	// 资源标识，比如资源类型为VMR，则该参数为vmrPkgId。
	TypeId *string `json:"typeId,omitempty"`

	// 资源标识对应的回显描述,比如资源类型为VMR，则该参数为vmrPkgName。
	TypeDesc *string `json:"typeDesc,omitempty"`

	// VMR模式。 - 0：个人会议ID - 1：云会议室 - 2：网络研讨会
	VmrMode *int32 `json:"vmrMode,omitempty"`

	// 资源数量。
	Count *int32 `json:"count,omitempty"`

	// 到期时间,utc时间戳。
	ExpireDate *int64 `json:"expireDate,omitempty"`

	// 资源对应的订单id。
	OrderId *string `json:"orderId,omitempty"`

	// 资源状态: - 0：正常 - 1：到期 - 2：停用
	Status *int32 `json:"status,omitempty"`

	// 标识资源是否可以编辑或删除。
	Editable *bool `json:"editable,omitempty"`
}

func (o QueryResourceResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResourceResultDto struct{}"
	}

	return strings.Join([]string{"QueryResourceResultDto", string(data)}, " ")
}
