package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTscDiagnoseResourcesReq struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`
}

func (o QueryTscDiagnoseResourcesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTscDiagnoseResourcesReq struct{}"
	}

	return strings.Join([]string{"QueryTscDiagnoseResourcesReq", string(data)}, " ")
}
