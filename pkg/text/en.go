package text

// Translation keys for English locale
const (
	// KeyAppTagline is the combined title/tagline
	KeyAppTagline = "app.tagline"
	// KeyWelcomeMessage is the welcome screen message
	KeyWelcomeMessage = "welcome.message"
	// KeyWelcomeFeature1 describes the first feature
	KeyWelcomeFeature1 = "welcome.feature1"
	// KeyWelcomeFeature2 describes the second feature
	KeyWelcomeFeature2 = "welcome.feature2"
	// KeyWelcomeFeature3 describes the third feature
	KeyWelcomeFeature3 = "welcome.feature3"
	// KeyWelcomeFeature4 describes the fourth feature
	KeyWelcomeFeature4 = "welcome.feature4"
	// KeyHelpQuit shows how to quit the application
	KeyHelpQuit = "help.quit"
	// KeyDockerStatusHealthy is shown when Docker connection is healthy
	KeyDockerStatusHealthy = "status.docker.healthy"
	// KeyDockerStatusDegraded is shown when Docker connection failed
	KeyDockerStatusDegraded = "status.docker.degraded"
	// KeyDockerStatusUnknown is shown before a health check runs
	KeyDockerStatusUnknown = "status.docker.unknown"
)

var translations = map[Locale]map[string]string{
	LocaleEN: {
		KeyWelcomeMessage: "This tool helps you identify and remove unused Docker containers, images,\n" +
			"volumes, and networks with confidence. Features include:",

		KeyAppTagline:      "docktidy - Spark joy in your Docker environment",
		KeyWelcomeFeature1: "Interactive resource selection with risk levels",
		KeyWelcomeFeature2: "Usage history tracking to protect active resources",
		KeyWelcomeFeature3: "Dry-run mode to preview changes before applying",
		KeyWelcomeFeature4: "Detailed cleanup history and recovery commands",

		KeyHelpQuit: "Press 'q', 'esc', or ctrl+c to quit",

		KeyDockerStatusHealthy:  "Docker: Connected to daemon",
		KeyDockerStatusDegraded: "Docker: Cannot reach daemon",
		KeyDockerStatusUnknown:  "Docker: Status check pending",
	},
}
