// Package text provides localized UI strings for docktidy.
//
// TODO: Locale switching is currently a placeholder for future i18n support.
// Only English (LocaleEN) is implemented. SetLocale/GetLocale exist for
// API completeness but adding new locales requires code changes.
package text

type Locale string

const (
	LocaleEN Locale = "en" // Only supported locale currently
)

var defaultLocale = LocaleEN

type Text struct {
	locale Locale // TODO: Locale switching not yet implemented
}

func New() *Text {
	return &Text{
		locale: defaultLocale,
	}
}

func Default() *Text {
	return New()
}

// SetLocale changes the locale for text retrieval.
// TODO: Only LocaleEN is supported. Other locales will fall back to English.
func (t *Text) SetLocale(locale Locale) {
	t.locale = locale
}

// GetLocale returns the current locale.
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
