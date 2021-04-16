package memory

// Memory implements storage.Storage interface
type Memory struct {
	memoryPoints
}

// New instance of Memory
func New() *Memory {
	return &Memory{}
}
