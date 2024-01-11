package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportLineageResponse Response Object
type ImportLineageResponse struct {

	// 血缘导入结果
	Body           *[]ObjectIdInfo `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ImportLineageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportLineageResponse struct{}"
	}

	return strings.Join([]string{"ImportLineageResponse", string(data)}, " ")
}
