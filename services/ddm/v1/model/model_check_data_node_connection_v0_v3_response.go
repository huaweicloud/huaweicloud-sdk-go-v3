package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDataNodeConnectionV0V3Response Response Object
type CheckDataNodeConnectionV0V3Response struct {

	// **参数解释**：  rds测试连通性相关信息的集合。  **参数范围**：  不涉及。
	RdsCheckInfos  *[]CheckRdsConnectionResqVo `json:"rds_check_infos,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o CheckDataNodeConnectionV0V3Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataNodeConnectionV0V3Response struct{}"
	}

	return strings.Join([]string{"CheckDataNodeConnectionV0V3Response", string(data)}, " ")
}
