package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteAkSkResponse struct {

	// 创建/删除的ak信息
	Ak             *string `json:"ak,omitempty" xml:"ak"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAkSkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAkSkResponse struct{}"
	}

	return strings.Join([]string{"DeleteAkSkResponse", string(data)}, " ")
}
