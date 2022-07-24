package loader

// Loader is the interface that defines the Loader service.
// The Loader service is responsible for loading transformed data
// to a destination db. It gets transformed data from a source chan.
type Loader interface {
	Load(sourceChan <-chan interface{}, errChan chan<- interface{})
}
