package config

var Constants = make(map[string]string)

func SetConstants() {
	Constants["BASE_URL"] = "http://localhost:9000/"
	Constants["STATIC_URL"] = "http://localhost:9000/public/"
}
