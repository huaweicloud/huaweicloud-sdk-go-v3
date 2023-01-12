package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// Response Object
type ShowBackupPathResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ShowBackupPathResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ShowBackupPathResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPathResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupPathResponse", string(data)}, " ")
}
