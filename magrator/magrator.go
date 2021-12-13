package magrator

type Magrator interface {
	Magrate()
}

func Magrate(driver interface{}) {
	if mag, ok := driver.(Magrator); ok {
		mag.Magrate()
	}
}
