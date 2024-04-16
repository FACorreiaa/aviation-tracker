// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package svg

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AirplaneIcon() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{"w-4 h-4"}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `app/svg/airplane.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" version=\"1.1\" id=\"Capa_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 229.345 229.345\" xml:space=\"preserve\"><g><g><g><path style=\"fill:#010002;\" d=\"M126.367,114.545h3.332l99.646,38.573v-32.72l-22.96-11.871V84.551\n				c0-2.99-2.433-5.413-5.442-5.413c-2.999,0-5.422,2.423-5.422,5.413v18.387l-29.271-15.114V62.959c0-2.98-2.443-5.383-5.432-5.383\n				s-5.422,2.403-5.422,5.383v19.267l-25.715-13.268h-1.378l1.798-41.699c-0.02-10.063-8.158-18.221-18.231-18.221\n				c-10.063,0-18.231,8.148-18.231,18.221l1.798,41.699h-2.198L74.243,79.431V62.959c0-2.98-2.433-5.383-5.432-5.383\n				c-2.999,0-5.403,2.403-5.403,5.383V85.44L39.54,98.591V80.829c0-2.99-2.443-5.422-5.422-5.422c-2.999,0-5.422,2.433-5.422,5.422\n				v23.742L0,120.398v32.73l93.227-38.573h4.143l2.697,62.998c0,3.615,0.244,6.986,1.006,9.839l-30.366,16.727l-3.879,16.189\n				l38.895-6.722h13.668l41.289,6.722l-3.879-16.189L122.83,186.6c0.606-2.687,0.84-5.764,0.84-9.057L126.367,114.545z\"></path></g></g></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
