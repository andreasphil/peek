package adapters

type FilePreviwer interface {
	ForFile(filename string) (string, error)
}
