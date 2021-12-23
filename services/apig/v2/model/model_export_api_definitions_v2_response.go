package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// Response Object
type ExportApiDefinitionsV2Response struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o ExportApiDefinitionsV2Response) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o ExportApiDefinitionsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportApiDefinitionsV2Response struct{}"
	}

	return strings.Join([]string{"ExportApiDefinitionsV2Response", string(data)}, " ")
}
