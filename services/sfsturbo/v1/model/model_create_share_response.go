package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateShareResponse struct {

	// 创建的SFS Turbo文件系统ID。
	Id *string `json:"id,omitempty"`

	// 创建的SFS Turbo文件系统名称。
	Name *string `json:"name,omitempty"`

	// SFS Turbo文件系统的状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShareResponse struct{}"
	}

	return strings.Join([]string{"CreateShareResponse", string(data)}, " ")
}
