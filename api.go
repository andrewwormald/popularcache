package popularcache

type API[T any] interface {
	Add(id string, item T)
	Collect(id string) (T, bool)
	List() []T
	TrimEnd(count int)
}
