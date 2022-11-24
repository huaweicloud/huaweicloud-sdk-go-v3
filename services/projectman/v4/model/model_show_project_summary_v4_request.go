package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowProjectSummaryV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`
}

func (o ShowProjectSummaryV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectSummaryV4Request struct{}"
	}

	return strings.Join([]string{"ShowProjectSummaryV4Request", string(data)}, " ")
}
