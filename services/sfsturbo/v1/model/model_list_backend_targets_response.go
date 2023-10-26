package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackendTargetsResponse Response Object
type ListBackendTargetsResponse struct {

	// 文件系统后端存储库列表元素个数
	Count *int32 `json:"count,omitempty"`

	// 文件系统后端存储库列表
	Targets *[]ShowBackendTargetInfoResponseBody `json:"targets,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBackendTargetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackendTargetsResponse struct{}"
	}

	return strings.Join([]string{"ListBackendTargetsResponse", string(data)}, " ")
}
