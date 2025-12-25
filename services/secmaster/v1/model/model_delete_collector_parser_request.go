package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCollectorParserRequest Request Object
type DeleteCollectorParserRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 解析器 ID
	ParserId string `json:"parser_id"`
}

func (o DeleteCollectorParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCollectorParserRequest struct{}"
	}

	return strings.Join([]string{"DeleteCollectorParserRequest", string(data)}, " ")
}
