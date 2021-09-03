package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/moooll/dogs-service/internal/errors"
	"github.com/moooll/dogs-service/internal/models"
	"github.com/moooll/dogs-service/internal/tokens"
	log "github.com/sirupsen/logrus"
)

// Tokens describes token pair
type Tokens struct {
	Access  string
	Refresh string
}

// Login godoc
// Login receives username-password pair and authenticates or registers if no user found
// @Summary Login or register a user
// @Description get credentials from a form and authenticate or register a user
// @ID login
// @Accept  x-www-form-urlencoded
// @Produce json
// @Success 200
// @Failure 401 {object} echo.HTTPError
// @Failure 401
// @Router /login [post]
func (s *Service) Login(c echo.Context) error {
	u := models.User{}
	u.Username = c.FormValue("username")
	u.Password = c.FormValue("password")
	vi, e := s.V.validate(u)
	if e != nil {
		log.Error("error creating user: "+fmt.Sprint(vi), e.Error())
		return echo.NewHTTPError(400, "error creating user, check the input")
	} else if vi != nil && e == nil {
		log.Error("error creating user: ", fmt.Sprint(vi))
		return echo.NewHTTPError(400, "error creating user, check the input")
	}

	err := s.St.Check(u.Username, u.Password)
	if err != nil {
		switch {
		case err.Error() == errors.AuthErrNotRegistered:
			u.ID, err = s.St.Register(u.Username, u.Password)
			if err != nil {
				log.Error("auth error: ", err.Error())
				return echo.NewHTTPError(401)
			}

			return c.NoContent(200)
		case err.Error() == errors.AuthErrWrongPwd:
			log.Error("auth error: ", err.Error())
			return echo.NewHTTPError(401, "error: wrong credentials")
		default:
			log.Error("auth error: ", err.Error())
			return echo.NewHTTPError(401)
		}
	}

	u.ID, err = s.St.GetUserByName(u.Username)
	if err != nil {
		log.Error("auth error: ", err.Error())
		return echo.NewHTTPError(500)
	}

	ts, er := s.generateTokenPair(u)
	if er != nil {
		log.Error("auth error: ", er.Error())
		return echo.NewHTTPError(401)
	}

	session := tokens.NewSession(u.ID, ts.Refresh)
	se := s.St.CreateSession(&session)
	if se != nil {
		log.Error("auth error: ", se.Error())
		return echo.NewHTTPError(500)
	}

	type res struct {
		Fingerprint uuid.UUID
		Tokens      Tokens
	}
	c.SetCookie(createRefreshCookie(ts.Refresh))
	return c.JSON(200, res{
		session.Fingerprint,
		ts})
}

// Refresh godoc
// Refresh refreshes token pair
// @Summary Refresh pair of tokens
// @Description get refresh token and fingerprint from client after token's validity check update existing session
// @ID refresh
// @Accept  x-www-form-urlencoded
// @Produce json
// @Success 200 {object} models.Tokens
// @Failure 401 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Router /refresh [post]
func (s *Service) Refresh(c echo.Context) error {
	type body struct {
		Refresh     uuid.UUID `json:"refresh"`
		Fingerprint uuid.UUID `json:"fingerprint"`
	}
	b := body{}
	err := c.Bind(&b)
	if err != nil {
		log.Error("refresh error: ", err.Error())
		return echo.NewHTTPError(500)
	}

	session, er := s.St.GetSession(b.Refresh)
	if er != nil {
		if er.Error() == "no rows in result set" {
			log.Error("refresh error: ", er.Error())
			return echo.NewHTTPError(401)
		}

		log.Error("refresh error: ", er.Error())
		return echo.NewHTTPError(500)
	}

	if session.Fingerprint != b.Fingerprint {
		log.Error("refresh error: ", "invalid refresh session")
		return echo.NewHTTPError(401, echo.Map{
			"error": "invalid refresh session",
		})
	}

	if session.ExpiresAt.Before(time.Now()) {
		log.Error("refresh error: ", "token expired")
		return echo.NewHTTPError(401, echo.Map{
			"error": "token expired",
		})
	}

	user, errr := s.St.GetUser(session.UserID)
	if errr != nil {
		log.Error("auth error: ", errr.Error())
		return echo.NewHTTPError(500)
	}

	ts, erro := s.generateTokenPair(user)
	if erro != nil {
		log.Error("refresh error: ", erro.Error())
		return echo.NewHTTPError(401)
	}

	newSession := tokens.NewSession(session.UserID, ts.Refresh)
	sesErr := s.St.UpdateSession(&newSession)
	if sesErr != nil {
		log.Error("refresh error: ", sesErr.Error())
		return echo.NewHTTPError(500)
	}

	c.SetCookie(createRefreshCookie(ts.Refresh))
	return c.JSON(200, ts)
}

func (s *Service) generateTokenPair(u models.User) (Tokens, error) {
	access, e := s.Sk.GenerateToken(u)
	if e != nil {
		log.Error("auth error: ", e.Error())
		return Tokens{}, e
	}

	refresh := tokens.NewResfreshToken()
	return Tokens{
		Access:  access,
		Refresh: refresh,
	}, nil
}

// Logout godoc
// Logout logouts the user and deletes existing session
// @Summary Logout user
// @Description get the refresh token from client and delete session
// @ID logout
// @Accept  x-www-form-urlencoded
// @Produce json
// @Success 401
// @Failure 500 {object} echo.HTTPError
// @Router /logout [post]
func (s *Service) Logout(c echo.Context) error {
	var refresh uuid.UUID
	err := c.Bind(&refresh)
	if err != nil {
		log.Error("logout error: ", err.Error())
		return echo.NewHTTPError(500)
	}

	er := s.St.DeleteSession(refresh)
	if er != nil {
		log.Error("logout error: ", er.Error())
		return echo.NewHTTPError(500)
	}

	return c.NoContent(401)
}

func createRefreshCookie(refresh string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "refreshToken"
	cookie.Value = refresh
	cookie.Expires = time.Now().Add(43200 * time.Minute)
	cookie.HttpOnly = true
	cookie.Path = "/auth"
	return cookie
}
