package singleton

var single *Singleton

func init() {
	single = new(Singleton)
}

func GetInstance() *Singleton {
	return single
}
