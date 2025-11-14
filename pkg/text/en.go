package text

// Translation keys for English locale
const (
	// KeyAppTagline is the combined title/tagline
	KeyAppTagline = "app.tagline"
	// KeyWelcomeMessage is the welcome screen message
	KeyWelcomeMessage = "welcome.message"
	// KeyWelcomeContinue instructs how to proceed from the welcome screen
	KeyWelcomeContinue = "welcome.continue"
	// KeyWelcomeFeature1 describes the first feature
	KeyWelcomeFeature1 = "welcome.feature1"
	// KeyWelcomeFeature2 describes the second feature
	KeyWelcomeFeature2 = "welcome.feature2"
	// KeyWelcomeFeature3 describes the third feature
	KeyWelcomeFeature3 = "welcome.feature3"
	// KeyWelcomeFeature4 describes the fourth feature
	KeyWelcomeFeature4 = "welcome.feature4"
	// KeyDashboardTitle labels the usage dashboard
	KeyDashboardTitle = "dashboard.title"
	// KeyDashboardBack instructs how to return to the welcome screen
	KeyDashboardBack = "dashboard.back"
	// KeyDashboardEmpty shows when usage data isn't available
	KeyDashboardEmpty = "dashboard.empty"
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
		KeyWelcomeContinue: "Press Enter to continue",
		KeyWelcomeFeature1: "Interactive resource selection with risk levels",
		KeyWelcomeFeature2: "Usage history tracking to protect active resources",
		KeyWelcomeFeature3: "Dry-run mode to preview changes before applying",
		KeyWelcomeFeature4: "Detailed cleanup history and recovery commands",
		KeyDashboardTitle:  "Docker Disk Usage",
		KeyDashboardBack:   "Press 'b' to return to the welcome menu",
		KeyDashboardEmpty:  "Disk usage data unavailable",

		KeyHelpQuit: "Press 'q', 'esc', or ctrl+c to quit",

		KeyDockerStatusHealthy:  "Docker: Connected to daemon",
		KeyDockerStatusDegraded: "Docker: Cannot reach daemon",
		KeyDockerStatusUnknown:  "Docker: Status check pending",
	},
}
