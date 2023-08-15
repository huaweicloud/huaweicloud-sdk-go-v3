package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ShowPointsResponse Response Object
type ShowPointsResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowPointsResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowPointsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPointsResponse struct{}"
	}

	return strings.Join([]string{"ShowPointsResponse", string(data)}, " ")
}
