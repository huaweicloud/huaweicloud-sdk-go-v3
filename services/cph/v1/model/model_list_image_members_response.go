package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageMembersResponse Response Object
type ListImageMembersResponse struct {

	// 镜像成员详情
	Members        *[]ListImageMembersView `json:"members,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListImageMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageMembersResponse struct{}"
	}

	return strings.Join([]string{"ListImageMembersResponse", string(data)}, " ")
}
