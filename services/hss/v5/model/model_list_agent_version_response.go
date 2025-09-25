package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentVersionResponse Response Object
type ListAgentVersionResponse struct {

	// agent版本信息详情列表
	DataList       *[]AgentVersionResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListAgentVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentVersionResponse struct{}"
	}

	return strings.Join([]string{"ListAgentVersionResponse", string(data)}, " ")
}
