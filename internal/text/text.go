package text

type Locale string

const (
	LocaleEN Locale = "en"
)

var defaultLocale = LocaleEN

type Text struct {
	locale Locale
}

func New() *Text {
	return &Text{
		locale: defaultLocale,
	}
}

func Default() *Text {
	return New()
}

func (t *Text) SetLocale(locale Locale) {
	t.locale = locale
}

func (t *Text) GetLocale() Locale {
	return t.locale
}

func (t *Text) Get(key string) string {
	messages, ok := translations[t.locale]
	if !ok {
		messages = translations[defaultLocale]
	}

	msg, ok := messages[key]
	if !ok {
		return key
	}

	return msg
}
