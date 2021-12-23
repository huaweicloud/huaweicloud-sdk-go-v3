package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadIssueImgResponse struct {
	// 图片id

	ImgId *string `json:"img_id,omitempty"`
	// 图片url v1改成v3作为下载图片请求

	ImgUrl         *string `json:"img_url,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadIssueImgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadIssueImgResponse struct{}"
	}

	return strings.Join([]string{"UploadIssueImgResponse", string(data)}, " ")
}
