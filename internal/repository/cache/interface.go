package cache

type Cache interface {
	Set(uuid string, value string)
	Get(uuid string) (string, bool)
}
