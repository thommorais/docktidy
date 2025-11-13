package text

import (
	"testing"
)

func TestNew(t *testing.T) {
	txt := New()
	if txt == nil {
		t.Fatal("New() returned nil")
	}

	if txt.locale != defaultLocale {
		t.Errorf("New() locale = %v, want %v", txt.locale, defaultLocale)
	}
}

func TestDefault(t *testing.T) {
	txt := Default()
	if txt == nil {
		t.Fatal("Default() returned nil")
	}

	if txt.locale != LocaleEN {
		t.Errorf("Default() locale = %v, want %v", txt.locale, LocaleEN)
	}
}

func TestSetLocale(t *testing.T) {
	txt := New()

	txt.SetLocale(LocaleEN)
	if txt.locale != LocaleEN {
		t.Errorf("SetLocale(LocaleEN) failed, got %v", txt.locale)
	}
}

func TestGetLocale(t *testing.T) {
	txt := New()

	if got := txt.GetLocale(); got != defaultLocale {
		t.Errorf("GetLocale() = %v, want %v", got, defaultLocale)
	}

	txt.SetLocale(LocaleEN)
	if got := txt.GetLocale(); got != LocaleEN {
		t.Errorf("GetLocale() after SetLocale = %v, want %v", got, LocaleEN)
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name   string
		locale Locale
		key    string
		want   string
	}{
		{
			name:   "existing key in English",
			locale: LocaleEN,
			key:    KeyAppTagline,
			want:   "docktidy - Spark joy in your Docker environment",
		},
		{
			name:   "help text key",
			locale: LocaleEN,
			key:    KeyHelpQuit,
			want:   "Press 'q', 'esc', or ctrl+c to quit",
		},
		{
			name:   "docker status key",
			locale: LocaleEN,
			key:    KeyDockerStatusHealthy,
			want:   "Docker: Connected to daemon",
		},
		{
			name:   "non-existent key returns key itself",
			locale: LocaleEN,
			key:    "nonexistent.key",
			want:   "nonexistent.key",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txt := New()
			txt.SetLocale(tt.locale)

			got := txt.Get(tt.key)
			if got != tt.want {
				t.Errorf("Get(%v) = %v, want %v", tt.key, got, tt.want)
			}
		})
	}
}

func TestGetWithInvalidLocale(t *testing.T) {
	txt := New()
	txt.SetLocale(Locale("invalid"))

	got := txt.Get(KeyAppTagline)
	want := "docktidy - Spark joy in your Docker environment"

	if got != want {
		t.Errorf("Get() with invalid locale should fallback to default, got %v, want %v", got, want)
	}
}

func TestAllKeysHaveTranslations(t *testing.T) {
	keys := []string{
		KeyAppTagline,
		KeyWelcomeMessage,
		KeyWelcomeFeature1,
		KeyWelcomeFeature2,
		KeyWelcomeFeature3,
		KeyWelcomeFeature4,
		KeyHelpQuit,
		KeyDockerStatusHealthy,
		KeyDockerStatusDegraded,
		KeyDockerStatusUnknown,
	}

	txt := Default()

	for _, key := range keys {
		got := txt.Get(key)
		if got == key {
			t.Errorf("Key %v has no translation", key)
		}
		if got == "" {
			t.Errorf("Key %v has empty translation", key)
		}
	}
}
