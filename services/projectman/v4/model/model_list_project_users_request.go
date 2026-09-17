package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectUsersRequest Request Object
type ListProjectUsersRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`
}

func (o ListProjectUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectUsersRequest struct{}"
	}

	return strings.Join([]string{"ListProjectUsersRequest", string(data)}, " ")
}
