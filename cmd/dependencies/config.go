package dependencies

type configuration struct {
	storage map[string]storage
}

type storage struct {
	memory map[int]string
}

var configs = map[string]configuration{
	"staging": {
		storage: map[string]storage{
			"number-repository": {
				memory: nil,
			},
		},
	},
	"production": {
		storage: map[string]storage{
			"number-repository": {
				memory: nil,
			},
		},
	},
}

//
// STORAGE
//

func initStorage(config storage) map[int]string {
	config.memory = make(map[int]string)
	return config.memory
}
