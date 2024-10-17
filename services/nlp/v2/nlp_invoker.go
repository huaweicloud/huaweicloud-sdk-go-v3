package v2

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/nlp/v2/model"
)

type RunAspectSentimentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunAspectSentimentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunAspectSentimentInvoker) Invoke() (*model.RunAspectSentimentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunAspectSentimentResponse), nil
	}
}

type RunAspectSentimentAdvanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunAspectSentimentAdvanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunAspectSentimentAdvanceInvoker) Invoke() (*model.RunAspectSentimentAdvanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunAspectSentimentAdvanceResponse), nil
	}
}

type RunClassificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunClassificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunClassificationInvoker) Invoke() (*model.RunClassificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunClassificationResponse), nil
	}
}

type RunConstituencyParserInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunConstituencyParserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunConstituencyParserInvoker) Invoke() (*model.RunConstituencyParserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunConstituencyParserResponse), nil
	}
}

type RunDependencyParserInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDependencyParserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunDependencyParserInvoker) Invoke() (*model.RunDependencyParserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDependencyParserResponse), nil
	}
}

type RunDocClassificationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDocClassificationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunDocClassificationInvoker) Invoke() (*model.RunDocClassificationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDocClassificationResponse), nil
	}
}

type RunDomainSentimentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunDomainSentimentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunDomainSentimentInvoker) Invoke() (*model.RunDomainSentimentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunDomainSentimentResponse), nil
	}
}

type RunEntityLinkingInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunEntityLinkingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunEntityLinkingInvoker) Invoke() (*model.RunEntityLinkingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunEntityLinkingResponse), nil
	}
}

type RunEntitySentimentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunEntitySentimentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunEntitySentimentInvoker) Invoke() (*model.RunEntitySentimentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunEntitySentimentResponse), nil
	}
}

type RunEventExtractionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunEventExtractionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunEventExtractionInvoker) Invoke() (*model.RunEventExtractionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunEventExtractionResponse), nil
	}
}

type RunFileTranslationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunFileTranslationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunFileTranslationInvoker) Invoke() (*model.RunFileTranslationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunFileTranslationResponse), nil
	}
}

type RunGetFileTranslationResultInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunGetFileTranslationResultInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunGetFileTranslationResultInvoker) Invoke() (*model.RunGetFileTranslationResultResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunGetFileTranslationResultResponse), nil
	}
}

type RunKeywordExtractInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunKeywordExtractInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunKeywordExtractInvoker) Invoke() (*model.RunKeywordExtractResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunKeywordExtractResponse), nil
	}
}

type RunLanguageDetectionInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunLanguageDetectionInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunLanguageDetectionInvoker) Invoke() (*model.RunLanguageDetectionResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunLanguageDetectionResponse), nil
	}
}

type RunMultiGrainedSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunMultiGrainedSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunMultiGrainedSegmentInvoker) Invoke() (*model.RunMultiGrainedSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunMultiGrainedSegmentResponse), nil
	}
}

type RunNerInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunNerInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunNerInvoker) Invoke() (*model.RunNerResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunNerResponse), nil
	}
}

type RunNerDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunNerDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunNerDomainInvoker) Invoke() (*model.RunNerDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunNerDomainResponse), nil
	}
}

type RunPoemInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunPoemInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunPoemInvoker) Invoke() (*model.RunPoemResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunPoemResponse), nil
	}
}

type RunSegmentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSegmentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSegmentInvoker) Invoke() (*model.RunSegmentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSegmentResponse), nil
	}
}

type RunSemanticParserInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSemanticParserInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSemanticParserInvoker) Invoke() (*model.RunSemanticParserResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSemanticParserResponse), nil
	}
}

type RunSentenceEmbeddingInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSentenceEmbeddingInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSentenceEmbeddingInvoker) Invoke() (*model.RunSentenceEmbeddingResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSentenceEmbeddingResponse), nil
	}
}

type RunSentimentInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSentimentInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSentimentInvoker) Invoke() (*model.RunSentimentResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSentimentResponse), nil
	}
}

type RunSummaryInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSummaryInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSummaryInvoker) Invoke() (*model.RunSummaryResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSummaryResponse), nil
	}
}

type RunSummaryDomainInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunSummaryDomainInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunSummaryDomainInvoker) Invoke() (*model.RunSummaryDomainResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunSummaryDomainResponse), nil
	}
}

type RunTextSimilarityInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunTextSimilarityInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunTextSimilarityInvoker) Invoke() (*model.RunTextSimilarityResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTextSimilarityResponse), nil
	}
}

type RunTextSimilarityAdvanceInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunTextSimilarityAdvanceInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunTextSimilarityAdvanceInvoker) Invoke() (*model.RunTextSimilarityAdvanceResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTextSimilarityAdvanceResponse), nil
	}
}

type RunTextTranslationInvoker struct {
	*invoker.BaseInvoker
}

func (i *RunTextTranslationInvoker) GetBaseInvoker() *invoker.BaseInvoker {
	return i.BaseInvoker
}

func (i *RunTextTranslationInvoker) Invoke() (*model.RunTextTranslationResponse, error) {
	if result, err := i.BaseInvoker.Invoke(); err != nil {
		return nil, err
	} else {
		return result.(*model.RunTextTranslationResponse), nil
	}
}
