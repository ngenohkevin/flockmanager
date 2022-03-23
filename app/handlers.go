package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/ngenohkevin/flock_manager/database"
	"github.com/ngenohkevin/flock_manager/models"
)

func (a *App) Initialize(config *database.Config) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require ",
		config.DB.Host,
		config.DB.Port,
		config.DB.Username,
		config.DB.Name,
		config.DB.Password,
	)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	a.Db = models.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.SetRouters()
}

//handlers for kuroiler

func (a *App) GetAllKuroilers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var kuroilers []models.Kuroiler
	db.Find(&kuroilers)
	err := a.writeJSON(w, http.StatusOK, kuroilers, "kuroiler")
	if err != nil {
		a.ErrorJSON(w, err)
	}
	return
}
func (a *App) CreateKuroiler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	kuroiler := models.Kuroiler{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&kuroiler); err != nil {
		a.ErrorJSON(w, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	if err := db.Save(&kuroiler).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusCreated, kuroiler, "kuroiler")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) GetKuroiler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	kuroiler := a.getKuroilerOr404(db, title, w, r)
	if kuroiler == nil {
		return
	}
	err := a.writeJSON(w, http.StatusOK, kuroiler, "kuroiler")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) UpdateKuroiler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	kuroiler := a.getKuroilerOr404(db, title, w, r)
	if kuroiler == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&kuroiler); err != nil {
		a.ErrorJSON(w, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	if err := db.Save(&kuroiler).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusOK, kuroiler, "kuroiler")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) DeleteKuroiler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	kuroiler := a.getKuroilerOr404(db, title, w, r)
	if kuroiler == nil {
		return
	}
	if err := db.Delete(&kuroiler).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusNoContent, nil, "kuroiler")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}

//handlers for RainbowRoosters

func (a *App) GetAllRainbowRoosters(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var rainbowrooster []models.RainbowRooster
	db.Find(&rainbowrooster)
	err := a.writeJSON(w, http.StatusOK, rainbowrooster, "rainbowrooster")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) CreateRainbowRooster(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	rainbowrooster := models.RainbowRooster{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rainbowrooster); err != nil {
		a.ErrorJSON(w, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	if err := db.Save(&rainbowrooster).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusCreated, rainbowrooster, "rainbowrooster")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) GetRainbowRooster(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	rainbowrooster := a.getRainbowRoosterOr404(db, title, w, r)
	if rainbowrooster == nil {
		return
	}
	err := a.writeJSON(w, http.StatusOK, rainbowrooster, "rainbowrooster")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) UpdateRainbowRooster(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	rainbowrooster := a.getRainbowRoosterOr404(db, title, w, r)
	if rainbowrooster == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rainbowrooster); err != nil {
		a.ErrorJSON(w, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	if err := db.Save(&rainbowrooster).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusOK, rainbowrooster, "rainbowrooster")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) DeleteRainbowRooster(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	rainbowrooster := a.getRainbowRoosterOr404(db, title, w, r)
	if rainbowrooster == nil {
		return
	}
	if err := db.Delete(&rainbowrooster).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusNoContent, nil, "rainbowrooster")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}

// handlers for Broilers

func (a *App) GetAllBroilers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var broilers []models.Broilers
	db.Find(&broilers)
	err := a.writeJSON(w, http.StatusOK, broilers, "broilers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) CreateBroiler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	broiler := models.Broilers{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&broiler); err != nil {
		a.ErrorJSON(w, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	if err := db.Save(&broiler).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusCreated, broiler, "broilers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) GetBroiler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	broiler := a.getBroilerOr404(db, title, w, r)
	if broiler == nil {
		return
	}
	err := a.writeJSON(w, http.StatusOK, broiler, "broilers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) UpdateBroiler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	broiler := a.getBroilerOr404(db, title, w, r)
	if broiler == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&broiler); err != nil {
		a.ErrorJSON(w, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)

	if err := db.Save(&broiler).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusOK, broiler, "broilers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) DeleteBroiler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	broiler := a.getBroilerOr404(db, title, w, r)
	if broiler == nil {
		return
	}
	if err := db.Delete(&broiler).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusNoContent, nil, "broilers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}

//Handlers for layers

func (a *App) GetAllLayers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var layers []models.Layers
	db.Find(&layers)
	err := a.writeJSON(w, http.StatusOK, layers, "layers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) CreateLayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	layer := models.Layers{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&layer); err != nil {
		a.ErrorJSON(w, err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	if err := db.Save(&layer).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusCreated, layer, "layers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) GetLayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	layer := a.getLayerOr404(db, title, w, r)
	if layer == nil {
		return
	}
	err := a.writeJSON(w, http.StatusOK, layer, "layers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) UpdateLayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	layer := a.getLayerOr404(db, title, w, r)
	if layer == nil {
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&layer); err != nil {
		a.ErrorJSON(w, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(r.Body)

	if err := db.Save(&layer).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusOK, layer, "layers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}
func (a *App) DeleteLayer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	layer := a.getLayerOr404(db, title, w, r)
	if layer == nil {
		return
	}
	if err := db.Delete(&layer).Error; err != nil {
		a.ErrorJSON(w, err)
		return
	}
	err := a.writeJSON(w, http.StatusNoContent, nil, "layers")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}

//handlers for flock

func (a *App) GetAllFlocks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var flocks []models.Flock
	db.Find(&flocks)
	err := a.writeJSON(w, http.StatusOK, flocks, "flock")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}

func (a *App) GetFlock(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	flocks := a.getFlockOr404(db, id, w, r)
	if flocks == nil {
		return
	}
	err := a.writeJSON(w, http.StatusOK, flocks, "flock")
	if err != nil {
		a.ErrorJSON(w, err)
	}
}

// func (a *App) CreateFlock(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
// 	flocks := models.Flock{}

// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&flocks); err != nil {
// 		a.ErrorJSON(w, err)
// 	}
// 	defer func(Body io.ReadCloser) {
// 		err := Body.Close()
// 		if err != nil {
// 			panic(err)
// 		}
// 	}(r.Body)
// 	if err := db.Save(&flocks).Error; err != nil {
// 		a.ErrorJSON(w, err)
// 		return
// 	}
// 	err := a.WriteJSON(w, http.StatusCreated, flocks)
// 	if err != nil {
// 		a.ErrorJSON(w, err)
// 	}
// }

// getKuroilerOr404 gets a kuroiler instance if exists, or respond the 404 error otherwise
func (a *App) getKuroilerOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *models.Kuroiler {
	kuroilers := models.Kuroiler{}
	if err := db.First(&kuroilers, models.Kuroiler{Title: title}).Error; err != nil {
		a.ErrorJSON(w, err)
		return nil
	}
	return &kuroilers
}

// getRainbowRoosterOr404 gets a rainbowrooster instance if exists, or respond the 404 error otherwise
func (a *App) getRainbowRoosterOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *models.RainbowRooster {
	roosters := models.RainbowRooster{}
	if err := db.First(&roosters, models.RainbowRooster{Title: title}).Error; err != nil {
		a.ErrorJSON(w, err)
		return nil
	}
	return &roosters
}

// getBroilerOr404 gets a Broiler instance if exists, or respond the 404 error otherwise
func (a *App) getBroilerOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *models.Broilers {
	broilers := models.Broilers{}
	if err := db.First(&broilers, models.Broilers{Title: title}).Error; err != nil {
		a.ErrorJSON(w, err)
		return nil
	}
	return &broilers
}

// getLayerOr404 gets a Layer instance if exists, or respond the 404 error otherwise
func (a *App) getLayerOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *models.Layers {
	layers := models.Layers{}
	if err := db.First(&layers, models.Layers{Title: title}).Error; err != nil {
		a.ErrorJSON(w, err)
		return nil
	}
	return &layers
}

//get Users or 404
func (a *App) getUser(db *gorm.DB, email string, w http.ResponseWriter, r *http.Request) *models.User {
	users := models.User{}
	if err := db.First(&users, models.User{Email: email}).Error; err != nil {
		//err = json.NewEncoder(w).Encode(err)
		//a.ErrorJSON(w, err)
		return nil
	}
	return &users
}

func (a *App) getFlockOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *models.Flock {
	flock := models.Flock{}
	if err := db.First(&flock, models.Flock{Id: id}).Error; err != nil {
		a.ErrorJSON(w, err)
		return nil
	}
	return &flock
}
