package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowMapTileResponse Response Object
type ShowMapTileResponse struct {
	ContentType    *string       `json:"Content-Type,omitempty"`
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowMapTileResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowMapTileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMapTileResponse struct{}"
	}

	return strings.Join([]string{"ShowMapTileResponse", string(data)}, " ")
}
