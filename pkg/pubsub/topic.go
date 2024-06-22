package pubsub

type Topic int

const (
	Sofia Topic = iota
	Prague
	Barcelona
	Rome
)

func (t Topic) String() string {
	switch t {
	case Sofia:
		return "Sofia"
	case Prague:
		return "Prague"
	case Barcelona:
		return "Barcelona"
	case Rome:
		return "Rome"
	default:
		return "Unknown"
	}
}