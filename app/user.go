package app

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/ngenohkevin/flock_manager/models"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const SecretKey = "secret"

func (a *App) Register(db *gorm.DB, w http.ResponseWriter, r *http.Request){
	var data map[string]string

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	users := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&users); err != nil{
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)

	if err := db.Save(&users).Error; err != nil{
		a.ErrorJSON(w, err)
		return
	}

	err := a.WriteJSON(w, http.StatusCreated,users)
	if err != nil{
		a.ErrorJSON(w, err)
	}
}

func (a *App)Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
	}
	var user models.User

	logins := a.getUser(db, data["email"], w , r)

	//a.Db.Where("email = ?", data["email"]).First(&user)
	js, err := json.Marshal(logins)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.ID == 0 {

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			err = json.NewEncoder(w).Encode(js)
		}



	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))
	//if errs != nil && errs == bcrypt.ErrMismatchedHashAndPassword {
	//	var  res = map[string]interface{}{"status": false, "message": "Invalid login credential. Please try again"}
	//	errs = json.NewEncoder(w).Encode(res)
	//	return
	//}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		res := "Error while generating token, Try again"
		err = json.NewEncoder(w).Encode(res)
		return
	}
	if err := a.WriteJSON(w, http.StatusOK, token); err != nil{
		resp := map[string]interface{}{"status": false, "message": "logged in"}
		if err != nil {
			err = json.NewEncoder(w).Encode(resp)
		}
	}
	//_, err = w.Write([]byte(token))
	//if err != nil {
	//	return
	//}
	//var resp = map[string]interface{}{"status": false, "message": "logged in"}
	//resp["token"] = token
	//if err := json.NewEncoder(w).Encode(resp); err != nil{
	//	return
	//}

	//http.SetCookie(w,
	//	&http.Cookie{
	//	Name:"token",
	//	Value: token,
	//	Expires: time.Now().Add(time.Hour * 24),
	//	})
}