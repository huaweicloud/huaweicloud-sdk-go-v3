package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepositoryResourcePermissionsResponse Response Object
type UpdateRepositoryResourcePermissionsResponse struct {

	// 返回状态码
	Status *int32 `json:"status,omitempty"`

	// 返回信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateRepositoryResourcePermissionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepositoryResourcePermissionsResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepositoryResourcePermissionsResponse", string(data)}, " ")
}
