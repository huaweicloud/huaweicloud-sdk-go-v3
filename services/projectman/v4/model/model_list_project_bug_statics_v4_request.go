package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectBugStaticsV4Request Request Object
type ListProjectBugStaticsV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`
}

func (o ListProjectBugStaticsV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectBugStaticsV4Request struct{}"
	}

	return strings.Join([]string{"ListProjectBugStaticsV4Request", string(data)}, " ")
}
