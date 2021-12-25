package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportLiveDataApiDefinitionsV2Response struct {
	// 导入成功信息

	Success *[]Success `json:"success,omitempty"`
	// 导入失败信息

	Failure *[]Failure `json:"failure,omitempty"`

	Swagger        *Swagger `json:"swagger,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ImportLiveDataApiDefinitionsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportLiveDataApiDefinitionsV2Response struct{}"
	}

	return strings.Join([]string{"ImportLiveDataApiDefinitionsV2Response", string(data)}, " ")
}
