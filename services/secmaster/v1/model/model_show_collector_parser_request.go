package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCollectorParserRequest Request Object
type ShowCollectorParserRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 解析器 ID
	ParserId string `json:"parser_id"`
}

func (o ShowCollectorParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCollectorParserRequest struct{}"
	}

	return strings.Join([]string{"ShowCollectorParserRequest", string(data)}, " ")
}
