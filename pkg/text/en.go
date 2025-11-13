package text

// Translation keys for English locale
const (
	// KeyAppTitle is the application title
	KeyAppTitle = "app.title"
	// KeyAppSubtitle is the application subtitle
	KeyAppSubtitle = "app.subtitle"
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
	// KeyWelcomePhilosophy describes the application philosophy
	KeyWelcomePhilosophy = "welcome.philosophy"
	// KeyHelpQuit shows how to quit the application
	KeyHelpQuit = "help.quit"
)

var translations = map[Locale]map[string]string{
	LocaleEN: {
		KeyAppTitle:    "docktidy - Docker Resource Manager",
		KeyAppSubtitle: "Spark joy in your Docker environment",

		KeyWelcomeMessage: "Welcome to docktidy! A TUI tool for safely cleaning up Docker resources.\n\n" +
			"This tool helps you identify and remove unused Docker containers, images,\n" +
			"volumes, and networks with confidence. Features include:",

		KeyWelcomeFeature1: "Interactive resource selection with risk levels",
		KeyWelcomeFeature2: "Usage history tracking to protect active resources",
		KeyWelcomeFeature3: "Dry-run mode to preview changes before applying",
		KeyWelcomeFeature4: "Detailed cleanup history and recovery commands",

		KeyWelcomePhilosophy: "Built with the \"spark joy\" philosophy - clean your Docker environment\n" +
			"with clarity and peace of mind.",

		KeyHelpQuit: "Press 'q', 'esc', or ctrl+c to quit",
	},
}
