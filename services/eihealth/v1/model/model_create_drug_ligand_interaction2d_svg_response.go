package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"io"

	"strings"
)

// CreateDrugLigandInteraction2dSvgResponse Response Object
type CreateDrugLigandInteraction2dSvgResponse struct {
	HttpStatusCode int           `json:"-"`
	Body           io.ReadCloser `json:"-" type:"stream"`
}

func (o CreateDrugLigandInteraction2dSvgResponse) Consume(writer io.Writer) (int64, error) {
	written, err := io.Copy(writer, o.Body)
	defer o.Body.Close()

	return written, err
}

func (o CreateDrugLigandInteraction2dSvgResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDrugLigandInteraction2dSvgResponse struct{}"
	}

	return strings.Join([]string{"CreateDrugLigandInteraction2dSvgResponse", string(data)}, " ")
}
