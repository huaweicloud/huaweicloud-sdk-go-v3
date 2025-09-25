package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecommendOfficialTemplateRequestBody 获取官方推荐模板请求体
type ListRecommendOfficialTemplateRequestBody struct {

	// 代码仓地址
	GitUrl string `json:"git_url"`

	// 代码仓分支名称你
	Branch string `json:"branch"`

	// 代码仓TAG
	Tags *string `json:"tags,omitempty"`
}

func (o ListRecommendOfficialTemplateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecommendOfficialTemplateRequestBody struct{}"
	}

	return strings.Join([]string{"ListRecommendOfficialTemplateRequestBody", string(data)}, " ")
}
