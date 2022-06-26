package config

type (
	Configuration struct {
		SqlDbSettings DbSettings `json:"SqlDbSettings"`
	}

	DbSettings struct {
		Uri          string `json:"Uri"`
		DatabaseName string `json:"DatabaseName"`
	}
)
