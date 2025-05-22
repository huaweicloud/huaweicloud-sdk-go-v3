package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// ExportDatabaseUsersResponse Response Object
type ExportDatabaseUsersResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportDatabaseUsersResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportDatabaseUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportDatabaseUsersResponse struct{}"
	}

	return strings.Join([]string{"ExportDatabaseUsersResponse", string(data)}, " ")
}
