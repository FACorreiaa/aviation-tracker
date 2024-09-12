// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package user

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/FACorreiaa/Aviation-tracker/app/models"
	"github.com/FACorreiaa/Aviation-tracker/app/view/components"
)

func LoginPage(login models.LoginPage) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"w-full bg-white\"><div class=\"mx-auto max-w-7xl\"><div class=\"flex flex-col lg:flex-row\"><div class=\"relative w-full bg-cover lg:w-6/12 xl:w-7/12 bg-gradient-to-r from-white via-white to-gray-100\"><div class=\"relative flex flex-col items-center justify-center w-full h-full px-10 my-20 lg:px-16 lg:my-0\"><div class=\"flex flex-col items-start space-y-8 tracking-tight lg:max-w-3xl\"><div class=\"relative\"><p class=\"mb-2 font-medium text-gray-700 uppercase\">Work smarter</p><h2 class=\"text-5xl font-bold text-gray-900 xl:text-6xl\">Features to help you work smarter</h2></div><p class=\"text-2xl text-gray-700\">We've created a simple formula to follow in order to gain more out of your business and your application.</p><a href=\"#_\" class=\"inline-block px-8 py-5 text-xl font-medium text-center text-white transition duration-200 bg-blue-600 rounded-lg hover:bg-blue-700 ease\" data-primary=\"blue-600\" data-rounded=\"rounded-lg\">Get Started Today</a></div></div></div><div class=\"w-full bg-white lg:w-6/12 xl:w-5/12\"><div class=\"flex flex-col items-start justify-start w-full h-full p-10 lg:p-16 xl:p-24\"><h4 class=\"w-full text-3xl font-bold\">Signin</h4><div class=\"relative w-full mt-10 space-y-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if login.Errors != nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"error-messages text-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, err := range login.Errors {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(err)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/view/user/login.templ`, Line: 31, Col: 19}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form method=\"post\"><fieldset><fieldset class=\"max-w-lg mx-auto mt-2\"><label for=\"email\" class=\"block text-sm font-medium text-gray-900 dark:text-white\">Email</label> <input class=\"block w-full px-4 py-4 mt-2 text-xl placeholder-gray-400 bg-gray-200 rounded-lg focus:outline-none focus:ring-4 focus:ring-blue-600 focus:ring-opacity-50\" data-primary=\"blue-600\" data-rounded=\"rounded-lg\" type=\"email\" placeholder=\"Email\" id=\"email\" name=\"email\" autocomplete=\"email\" required></fieldset><fieldset class=\"max-w-lg mx-auto mt-2\"><label for=\"password\" class=\"block text-sm font-medium text-gray-900 dark:text-white\">Password</label> <input class=\"block w-full px-4 py-4 mt-2 text-xl placeholder-gray-400 bg-gray-200 rounded-lg focus:outline-none focus:ring-4 focus:ring-blue-600 focus:ring-opacity-50\" data-primary=\"blue-600\" data-rounded=\"rounded-lg\" type=\"password\" placeholder=\"Password\" name=\"password\" autocomplete=\"current-password\" required></fieldset></fieldset><div class=\"flex items-center justify-center max-w-lg mx-auto mt-2\"><button type=\"submit\" class=\"focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-green-600 dark:hover:bg-green-700 dark:focus:ring-green-800\">Sign in</button><p class=\"text-sm px-5 py-2.5 me-2 mb-2 text-center\"><a href=\"/register\">Need an account?</a></p></div><div class=\"flex items-center justify-center max-w-lg mx-auto mt-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.GoogleSignin("Sign in with Google").Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></form></div></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
