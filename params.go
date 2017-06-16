package mapgen

// Params contains raw template paramters
type Params struct {
	Package    string
	Exported   bool
	UseRWMutex bool
	MapName    string
	KeyType    string
	ValType    string
}
