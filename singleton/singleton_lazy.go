package _singleton

import "sync"

var (
	lazyInstance *Singleton
	once         sync.Once
)

func GetLazySingleton() *Singleton {
	if lazyInstance == nil {
		once.Do(func() {
			lazyInstance = &Singleton{}
		})
	}
	return lazyInstance
}
