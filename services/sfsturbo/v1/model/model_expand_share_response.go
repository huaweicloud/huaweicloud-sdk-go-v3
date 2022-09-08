package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExpandShareResponse struct {

	// SFS Turbo文件系统ID。
	Id *string `json:"id,omitempty"`

	// SFS Turbo文件系统名称。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandShareResponse struct{}"
	}

	return strings.Join([]string{"ExpandShareResponse", string(data)}, " ")
}
