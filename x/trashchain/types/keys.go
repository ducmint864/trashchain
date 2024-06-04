package types

const (
	// ModuleName defines the module name
	ModuleName = "trashchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_trashchain"
)

var (
	ParamsKey = []byte("p_trashchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
