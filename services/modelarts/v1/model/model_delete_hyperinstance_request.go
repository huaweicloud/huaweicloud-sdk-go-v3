package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHyperinstanceRequest Request Object
type DeleteHyperinstanceRequest struct {

	// **参数解释**：Lite Server 超节点ID。
	Id string `json:"id"`
}

func (o DeleteHyperinstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHyperinstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteHyperinstanceRequest", string(data)}, " ")
}
