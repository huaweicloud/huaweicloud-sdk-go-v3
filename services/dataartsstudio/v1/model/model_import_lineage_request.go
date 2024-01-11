package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportLineageRequest Request Object
type ImportLineageRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 血缘信息列表
	Body *[]TableLineage `json:"body,omitempty"`
}

func (o ImportLineageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportLineageRequest struct{}"
	}

	return strings.Join([]string{"ImportLineageRequest", string(data)}, " ")
}
