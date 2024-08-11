package _singleton

type Singleton struct {
	_ string
}

var instance *Singleton

func init() {
	instance = &Singleton{}
}

func getSingleton() *Singleton {
	return instance
}
