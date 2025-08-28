package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePolicyGroupRequestInfo 需要删除的策略组ID列表信息
type DeletePolicyGroupRequestInfo struct {

	// **参数解释**: 需要被删除的策略组的策略组ID列表，仅支持删除非默认并且未关联服务器的策略组 **约束限制**: 需要使用 ListPolicyGroup 接口查询旗舰版和容器版策略组，在 ListPolicyGroup 接口的响应体中，deletable 等于 true 的 group_id 是可以被删除的策略组ID **取值范围**: 最少1条，最多50条 **默认取值**: 不涉及
	IdList []string `json:"id_list"`
}

func (o DeletePolicyGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePolicyGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"DeletePolicyGroupRequestInfo", string(data)}, " ")
}
