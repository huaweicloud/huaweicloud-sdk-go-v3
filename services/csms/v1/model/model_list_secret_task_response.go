package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecretTaskResponse Response Object
type ListSecretTaskResponse struct {

	// 任务数量
	Total *int32 `json:"total,omitempty"`

	// 凭据任务列表。
	Tasks *[]SecretTask `json:"tasks,omitempty"`

	// 下一页查询地址（本页的末尾任务ID）。
	NextMarker     *string `json:"next_marker,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSecretTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretTaskResponse struct{}"
	}

	return strings.Join([]string{"ListSecretTaskResponse", string(data)}, " ")
}
