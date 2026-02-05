package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckDataNodeConnectionResponse Response Object
type CheckDataNodeConnectionResponse struct {

	// **参数解释**：  rds测试连通性相关信息的集合。  **参数范围**：  不涉及。
	RdsCheckInfos  *[]CheckRdsConnectionResqVo `json:"rds_check_infos,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o CheckDataNodeConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataNodeConnectionResponse struct{}"
	}

	return strings.Join([]string{"CheckDataNodeConnectionResponse", string(data)}, " ")
}
