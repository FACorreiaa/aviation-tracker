package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/FACorreiaa/Aviation-tracker/app/view/user"

	"github.com/FACorreiaa/Aviation-tracker/app/models"
	"github.com/FACorreiaa/Aviation-tracker/core/account"
)

func (h *Handlers) registerPage(w http.ResponseWriter, r *http.Request) error {
	register := user.RegisterPage(models.RegisterPage{})
	return h.CreateLayout(w, r, "Sign up", register).Render(context.Background(), w)
}

func (h *Handlers) registerPost(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	var form account.RegisterForm
	var err error

	var token *account.Token
	err = h.formDecoder.Decode(&form, r.PostForm)
	if err == nil {
		token, err = h.core.accounts.RegisterNewAccount(r.Context(), form)
	}

	if err != nil {
		register := user.RegisterPage(models.RegisterPage{Errors: h.formErrors(err)})
		return h.CreateLayout(w, r, "Sign up", register).Render(context.Background(), w)
	}

	session, _ := h.sessions.Get(r, "auth")
	session.Values["token"] = token
	if err := session.Save(r, w); err != nil {
		return errors.New("failed to save session")
	}

	http.Redirect(w, r, "/settings", http.StatusSeeOther)
	return nil
}
