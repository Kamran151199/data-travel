package extractor

// Extractor is an interface that defines the extractor service.
// The extractor service is responsible for extracting data from a source db and
// sending it to a sink chan for transformation.
type Extractor interface {
	Extract(sinkChan chan<- interface{}, errChan chan<- interface{})
}
