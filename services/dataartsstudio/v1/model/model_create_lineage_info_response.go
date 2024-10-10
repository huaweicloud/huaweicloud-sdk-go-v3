package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLineageInfoResponse Response Object
type CreateLineageInfoResponse struct {

	// 血缘导入结果
	Body           *[]ObjectIdInfo `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateLineageInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLineageInfoResponse struct{}"
	}

	return strings.Join([]string{"CreateLineageInfoResponse", string(data)}, " ")
}
