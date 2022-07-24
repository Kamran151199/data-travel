package transformer

// Transformer is interface that defines transformer service.
// The transformer service is responsible for transforming data from one format
// to another. It gets data from a source chan and sends it to a sink chan.
type Transformer interface {
	Transform(sourceChan <-chan interface{}, sinkChan chan<- interface{}, errChan chan<- interface{})
}
