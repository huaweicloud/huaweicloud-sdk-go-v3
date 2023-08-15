package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareFaceByUrlResponse Response Object
type CompareFaceByUrlResponse struct {
	Image1Face *CompareFace `json:"image1_face,omitempty"`

	Image2Face *CompareFace `json:"image2_face,omitempty"`

	// 人脸相似度，1表示最大，0表示最小，值越大表示越相似。一般情况下超过0.93即可认为是同一个人。 调用失败时无此字段。
	Similarity     *float64 `json:"similarity,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o CompareFaceByUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareFaceByUrlResponse struct{}"
	}

	return strings.Join([]string{"CompareFaceByUrlResponse", string(data)}, " ")
}
