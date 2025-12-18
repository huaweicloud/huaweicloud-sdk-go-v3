package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListParameterGroupApplyHistoryV0V3Response Response Object
type ListParameterGroupApplyHistoryV0V3Response struct {

	// **参数解释**：  参数组应用记录相关信息的集合。  **参数范围**：  不涉及。
	RdsCheckInfos  *[]ApplyHistory `json:"rds_check_infos,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListParameterGroupApplyHistoryV0V3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListParameterGroupApplyHistoryV0V3Response struct{}"
	}

	return strings.Join([]string{"ListParameterGroupApplyHistoryV0V3Response", string(data)}, " ")
}
