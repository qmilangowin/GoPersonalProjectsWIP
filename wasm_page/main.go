package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var web fs.FS

type resourcesFS struct {
	http.Handler
}

func newReourcesFS(fsys fs.FS) app.ResourceProvider {
	return resourcesFS{
		Handler: http.FileServer((http.FS(fsys))),
	}
}

func (resourcesFS) Package() string { return "" }
func (resourcesFS) Static() string  { return "" }
func (resourcesFS) AppWASM() string { return "/web/app.wasm" }

type LoginPage struct {
	app.Compo
}

func (p *LoginPage) Render() app.UI {
	return app.Div().Class("min-h-screen flex items-center justify-center bg-gray-100").Body(
		app.Div().Class("max-w-md w-full space-y-8").Body(
			app.Div().Class("bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4").Body(
				app.H2().Class("mb-6 text-center text-3xl font-extrabold text-gray-900").Text("Sign in to your account"),
				app.Form().Class("mt-8 space-y-6").Body(
					app.Div().Class("rounded-md shadow-sm -space-y-px").Body(
						app.Div().Class("mb-3").Body(
							app.Label().Class("block text-gray-700 text-sm font-bold mb-2").For("email").Text("Email"),
							app.Input().Type("email").ID("email").Class("appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm").Placeholder("Email address"),
						),
						app.Div().Body(
							app.Label().Class("block text-gray-700 text-sm font-bold mb-2").For("password").Text("Password"),
							app.Input().Type("password").ID("password").Class("appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm").Placeholder("Password"),
						),
					),
					app.Div().Class("flex items-center justify-between").Body(
						app.Button().Class("group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500").Type("button").Text("Sign in"),
					),
				),
			),
		),
	)
}

//==

// type SignInBasic struct {
// 	app.Compo
// 	RememberMe bool
// }

// func (s *SignInBasic) Render() app.UI {
// 	return app.Div().Class("ui middle aligned center aligned grid").Style("height", "100vh").Body(
// 		app.Div().Class("column").Style("max-width", "450px").Body(
// 			// This div acts as a card container
// 			app.Div().Class("ui card").Body(
// 				app.Div().Class("content").Body(
// 					app.H4().Class("ui teal image header").Text("Sign in"),
// 					// Your form and other elements go here
// 					app.Div().Class("ui form").Body(
// 						app.Div().Class("field").Body(
// 							app.Input().Type("email").Placeholder("Email"),
// 						),
// 						app.Div().Class("field").Body(
// 							app.Input().Type("password").Placeholder("Password"),
// 						),
// 						// Add more form elements as needed
// 					),
// 					app.Button().Class("ui fluid large teal submit button").Text("Login"),
// 					// Other card content...
// 				),
// 				// Optionally add more divs for additional card sections
// 			),
// 		),
// 	)
// }

// // OnRememberMeChange handles the change event for the "Remember Me" switch.
// func (s *SignInBasic) OnRememberMeChange(ctx app.Context, e app.Event) {
// 	s.RememberMe = !s.RememberMe
// 	s.Update()
// }

// // OnSignIn handles the sign-in button click.
// func (s *SignInBasic) OnSignIn(ctx app.Context, e app.Event) {
// 	fmt.Println("SIGNING IN HER")
// }

func main() {
	h := LoginPage{}
	app.Route("/", &h)
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Admin",
		Description: "Admin portal",
		Styles: []string{
			"/web/semantic.css",
			"https://cdnjs.cloudflare.com/ajax/libs/semantic-ui/2.4.0/semantic.min.css",
			"https://cdn.tailwindcss.com",
		},
		// Scripts: []string{
		// 	"https://cdn.muicss.com/mui-0.10.3/js/mui.min.js",
		// },
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
