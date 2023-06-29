package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FrozenInfo 云服务或资源实例冻结信息。
type FrozenInfo struct {

	// 云服务或资源实例状态，取值： - 0：解冻/正常（云服务恢复正常）。 - 1：冻结（资源和数据会保留，但租户无法再正常使用云服务）。 - 2：删除/终止（资源和数据将清除）。
	Status *int32 `json:"status,omitempty"`

	// 在冻结/解冻操作下，取值： - 1（默认值）：冻结可释放。 - 2：冻结不可释放。 - 3：冻结后不可续费。
	Effect *int32 `json:"effect,omitempty"`

	// 更新云服务状态的业务场景列表，取值： - ARREAR（默认值）：欠费场景。为正常的运营业务场景，包括包周期资源到期、按需资源扣费失败。 - POLICE：公安冻结场景。 - ILLEGAL：违规冻结场景。 - VERIFY：客户未实名认证冻结场景。 - PARTNER：合作伙伴冻结（合作伙伴冻结子客户资源）。
	Scene *[]string `json:"scene,omitempty"`
}

func (o FrozenInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrozenInfo struct{}"
	}

	return strings.Join([]string{"FrozenInfo", string(data)}, " ")
}
