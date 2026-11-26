package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowAssetTagsInfoResponse Response Object
type ShowAssetTagsInfoResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowAssetTagsInfoResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowAssetTagsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetTagsInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetTagsInfoResponse", string(data)}, " ")
}
