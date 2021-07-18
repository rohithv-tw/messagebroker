package Config

type IConfig interface {
	GetHost() string
	GetTopic() string
}
