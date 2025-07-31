package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCompareProgressResponse Response Object
type ShowCompareProgressResponse struct {
	FullInfo *QueryCompareJobProgressRespFullInfo `json:"full_info,omitempty"`

	IncreInfo *QueryCompareJobProgressRespIncreInfo `json:"incre_info,omitempty"`

	GlobalInfo     *QueryCompareJobProgressRespGlobalInfo `json:"global_info,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ShowCompareProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompareProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowCompareProgressResponse", string(data)}, " ")
}
