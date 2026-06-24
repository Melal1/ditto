package keyboard

type Standard int

const (
	ANSI Standard = iota
	ISO
)

func (s Standard) String() string {
	switch s {
	case ISO:
		return "iso"
	default:
		return "ansi"
	}
}

func (s Standard) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *Standard) UnmarshalText(text []byte) error {
	switch string(text) {
	case "iso":
		*s = ISO
	default:
		*s = ANSI
	}
	return nil
}
