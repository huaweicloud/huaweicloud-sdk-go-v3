package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProjectDemandStaticV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`
}

func (o ListProjectDemandStaticV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectDemandStaticV4Request struct{}"
	}

	return strings.Join([]string{"ListProjectDemandStaticV4Request", string(data)}, " ")
}
