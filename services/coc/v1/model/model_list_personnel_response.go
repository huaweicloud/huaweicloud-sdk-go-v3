package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPersonnelResponse Response Object
type ListPersonnelResponse struct {

	// 人员信息
	Personnel *[]PersonnelResponse `json:"personnel,omitempty"`

	// 总数
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPersonnelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPersonnelResponse struct{}"
	}

	return strings.Join([]string{"ListPersonnelResponse", string(data)}, " ")
}
