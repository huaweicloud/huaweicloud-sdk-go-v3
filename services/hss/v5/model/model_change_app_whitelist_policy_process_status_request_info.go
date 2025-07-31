package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAppWhitelistPolicyProcessStatusRequestInfo 标记策略识别进程
type ChangeAppWhitelistPolicyProcessStatusRequestInfo struct {

	// **参数解释**： 进程可信状态 **约束限制**: 不涉及 **取值范围**: - trust：可信 - suspicious：可疑 - malicious：恶意 - unknown：未知  **默认取值**: 不涉及
	ProcessStatus string `json:"process_status"`

	// 进程hash列表
	ProcessHashList []string `json:"process_hash_list"`
}

func (o ChangeAppWhitelistPolicyProcessStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAppWhitelistPolicyProcessStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeAppWhitelistPolicyProcessStatusRequestInfo", string(data)}, " ")
}
