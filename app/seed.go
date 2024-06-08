package app

import (
	"math/rand"

	"github.com/andrewarrow/feedback/markup"
	"github.com/andrewarrow/feedback/router"
)

func Seed(c *router.Context) {
	for _, item := range projects {
		c.Params = map[string]any{}
		c.Params["name"] = item
		c.Params["color"] = markup.RandomColor()
		c.ValidateCreate("project")
		c.Insert("project")
		one := c.One("project", "where name=$1", item)
		if len(one) == 0 {
			continue
		}
		for i := 0; i < 9; i++ {
			c.Params = map[string]any{}
			c.Params["name"] = tasks[rand.Intn(len(tasks))-1]
			if rand.Intn(100) > 20 {
				c.Params["project_id"] = one["id"]
			}
			c.ValidateCreate("task")
			c.Insert("task")
		}
	}
}

var projects = []string{
	"SnapSavvy",
	"MindMate",
	"FitTrack",
	"TaskMaster",
	"RecipeGuru",
	"TravelBuddy",
	"WalletWise",
	"LearnSwift",
	"MusicMix",
	"HomeHaven",
}

var tasks = []string{
	"Implement a search feature",
	"Save the user's session data",
	"Design the user interface",
	"Create a login page",
	"Develop a registration form",
	"Set up user authentication",
	"Create a user profile page",
	"Integrate social media login options",
	"Implement password recovery",
	"Develop an admin dashboard",
	"Set up a database",
	"Optimize database queries",
	"Create a RESTful API",
	"Implement data validation",
	"Set up error handling",
	"Develop a notification system",
	"Create a messaging system",
	"Integrate third-party APIs",
	"Implement file uploads",
	"Set up image processing",
	"Create a reporting feature",
	"Develop a user settings page",
	"Implement a payment gateway",
	"Create a subscription system",
	"Set up email notifications",
	"Develop a mobile-responsive design",
	"Implement push notifications",
	"Create a comment system",
	"Set up user permissions and roles",
	"Develop a search engine optimization (SEO) strategy",
	"Create a sitemap",
	"Implement analytics tracking",
	"Optimize page load times",
	"Set up server-side caching",
	"Develop a shopping cart",
	"Create a product catalog",
	"Implement a rating and review system",
	"Set up a blog section",
	"Develop a contact form",
	"Implement live chat support",
	"Create a FAQ section",
	"Develop a tutorial or onboarding process",
	"Implement a content management system (CMS)",
	"Set up a version control system",
	"Create a staging environment",
	"Develop unit tests",
	"Implement integration tests",
	"Create a backup system",
	"Set up continuous integration/continuous deployment (CI/CD)",
	"Implement security measures",
	"Develop a logout feature",
	"Create a password change option",
	"Set up role-based access control",
	"Implement data encryption",
	"Develop a user activity log",
	"Create a user feedback system",
	"Implement a referral program",
	"Set up A/B testing",
	"Develop a user survey feature",
	"Implement multi-language support",
	"Create a terms and conditions page",
	"Set up a privacy policy page",
	"Develop a cookie consent banner",
	"Implement session timeout",
	"Create a custom error page",
	"Develop a progress tracker",
	"Set up automated emails",
	"Implement a drag-and-drop interface",
	"Create a calendar integration",
	"Develop a search filter",
	"Implement pagination",
	"Set up infinite scrolling",
	"Create a user activity feed",
	"Develop a file download feature",
	"Implement offline access",
	"Create a print-friendly version",
	"Set up a dark mode option",
	"Develop a bookmarking feature",
	"Implement a recommendation engine",
	"Create a wishlist feature",
	"Set up social sharing options",
	"Develop an import/export feature",
	"Implement two-factor authentication",
	"Create an API documentation",
	"Set up user analytics",
	"Develop a user segmentation feature",
	"Implement a coupon system",
	"Create a transaction history page",
	"Set up a loyalty program",
	"Develop a geolocation feature",
	"Implement a voice search",
	"Create a user onboarding checklist",
	"Set up real-time updates",
	"Develop a heatmap for user activity",
	"Implement custom user avatars",
	"Create a markdown editor",
	"Set up a subscription reminder",
	"Develop a chatbot",
	"Implement a dynamic form builder",
	"Create a multi-step form",
	"Set up a service status page",
	"Develop an API rate limiting",
	"Implement cross-platform synchronization",
}
