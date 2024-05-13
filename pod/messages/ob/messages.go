package ob

type Args map[string]string

type Recipient interface {
	SendTo() string
	SendName() []string
}
