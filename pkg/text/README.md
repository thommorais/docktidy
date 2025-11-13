# Text Package

This package manages all UI text and messages for docktidy, with infrastructure for future internationalization (i18n).

**Note:** Locale switching is currently a TODO. Only English is implemented. SetLocale/GetLocale methods exist for API design but adding new locales requires code changes.

## Structure

```md
pkg/text/
├── text.go       # Core text management and locale switching
├── en.go         # English translations
└── README.md     # This file
```

## Usage

```go
import "github.com/thommorais/docktidy/pkg/text"

// Create text instance with default locale (English)
t := text.Default()

// Get a translated string
title := t.Get(text.KeyAppTitle)

// TODO: Locale switching not yet implemented
// t.SetLocale(text.LocaleES) // Would fall back to English
```

## Adding New Text

1. Add a constant key in the appropriate locale file (e.g., `en.go`):

```go
const (
    KeyMyNewText = "my.new.text"
)
```

2. Add the translation to the `translations` map:

```go
var translations = map[Locale]map[string]string{
    LocaleEN: {
        KeyMyNewText: "My new text in English",
    },
}
```

3. Use the key in your code:

```go
message := t.Get(text.KeyMyNewText)
```

## Adding a New Locale (TODO - Future Feature)

**Note:** This is currently not implemented. The steps below describe the intended design.

1. Create a new file (e.g., `es.go` for Spanish)
2. Add the locale constant to `text.go`:

```go
const (
    LocaleEN Locale = "en"
    LocaleES Locale = "es" // TODO: Not yet implemented
)
```

3. Add translations to the `translations` map in the new file:

```go
var spanishTranslations = map[string]string{
    KeyAppTitle: "docktidy - Gestor de Recursos Docker",
    // ... more translations
}

func init() {
    translations[LocaleES] = spanishTranslations
}
```

## Key Naming Convention

Use dot notation with the pattern: `<section>.<subsection>.<item>`

Examples:

- `app.title` - Application-level text
- `welcome.message` - Welcome screen text
- `help.quit` - Help text for quitting
- `menu.prune` - Menu items
- `error.docker_connection` - Error messages

## Design Principles

- All user-facing text must go through this package
- No hardcoded strings in UI code
- Keys should be descriptive and hierarchical
- Fallback to English if translation missing
- Keep translations in separate files per locale
