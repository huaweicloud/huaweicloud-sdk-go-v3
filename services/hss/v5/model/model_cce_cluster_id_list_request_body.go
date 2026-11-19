package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CceClusterIdListRequestBody CCE集群配置请求
type CceClusterIdListRequestBody struct {

	// **参数解释**： 集群ID列表 **约束限制**： 不涉及 **取值范围**： 数组长度0-200，每个元素字符长度0-128位 **默认取值**： 不涉及
	ClusterIdList []string `json:"cluster_id_list"`

	// **参数解释**： 查询类型 **约束限制**： 不涉及 **取值范围**： - image：镜像风险 - baseline：基线风险 - vul：漏洞风险 - event：入侵风险  **默认取值**： 不涉及
	DetectType *string `json:"detect_type,omitempty"`
}

func (o CceClusterIdListRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CceClusterIdListRequestBody struct{}"
	}

	return strings.Join([]string{"CceClusterIdListRequestBody", string(data)}, " ")
}
