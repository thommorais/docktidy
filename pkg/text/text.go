// Package text provides localized UI strings for docktidy.
//
// TODO: Locale switching is currently a placeholder for future i18n support.
// Only English (LocaleEN) is implemented. SetLocale/GetLocale exist for
// API completeness but adding new locales requires code changes.
package text

// Locale represents a language/region combination for translations.
type Locale string

const (
	// LocaleEN is the English locale (currently the only supported locale).
	LocaleEN Locale = "en"
)

var defaultLocale = LocaleEN

// Text manages localized strings for the UI.
type Text struct {
	locale Locale // TODO: Locale switching not yet implemented
}

// New creates a new Text instance with the default locale.
func New() *Text {
	return &Text{
		locale: defaultLocale,
	}
}

// Default creates a new Text instance with the default locale (alias for New).
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

// Get retrieves a translated string by key, falling back to the key itself if not found.
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
