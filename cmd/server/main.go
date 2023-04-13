package main

import (
	"fmt"
	"log"
	"os"
	"time"
	_ "time/tzdata"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rows15/CTDGolangBackendIII/cmd/server/handler"
	"github.com/rows15/CTDGolangBackendIII/docs"
	"github.com/rows15/CTDGolangBackendIII/internal/appointment"
	"github.com/rows15/CTDGolangBackendIII/internal/dentist"
	"github.com/rows15/CTDGolangBackendIII/internal/patient"
	"github.com/rows15/CTDGolangBackendIII/pkg/middleware"
	"github.com/rows15/CTDGolangBackendIII/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

}

// @title Dental Clinic API
// @version 1.0

// @contact.name Marcelo Ramos
// @contact.url https://github.com/rows15

// @description This API handle appointments, dentists and patients for dental clinic system.

// @termsOfService https://www.linkedin.com/in/rows15

// @license.name Apache 2.0
// @license.url https://www.apache.org/license/LICENSE-2.0.html

// @securityDefinitions.apikey SECRET_TOKEN
// @in header
// @name SECRET_TOKEN

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file", err.Error())
	}

	// 	DB INITIALIZATION

	sqlStore := store.NewSQLStore()
	apStore := store.NewSQLAp()
	//Handlers INITIALIZATION

	appRepo := appointment.NewRepository(apStore)
	appService := appointment.NewService(appRepo)
	appHandler := handler.NewAppointmentHandler(appService)

	dentistRepo := dentist.NewRepository(sqlStore)
	dentistService := dentist.NewService(dentistRepo)

	dentistHandler := handler.NewDentistHandler(dentistService)

	patientRepo := patient.NewRepository(sqlStore)
	patientService := patient.NewService(patientRepo)
	patientHandler := handler.NewPatientHandler(patientService)

	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.BasePath = os.Getenv("BASE_PATH")
	r.GET("/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.Authentication())

	api := r.Group("/api/v1")
	{
		appointments := api.Group("/appointments")
		{
			appointments.GET("", appHandler.GetAll())
			appointments.GET(":id", appHandler.GetByID())
			appointments.GET("/patient/:identity_number", appHandler.GetAllByIdentityNumber())
			appointments.GET("/dentist/:license_number", appHandler.GetAllByLicenseNumber())

			appointments.POST("", appHandler.Post())
			appointments.PUT(":id", appHandler.Put())
			appointments.PATCH(":id", appHandler.Patch())
			appointments.DELETE(":id", appHandler.Delete())
		}
		dentists := api.Group("/dentists")
		{
			dentists.GET("", dentistHandler.GetAll())
			dentists.GET(":id", dentistHandler.GetByID())

			dentists.POST("", dentistHandler.Post())
			dentists.PUT(":id", dentistHandler.Put())
			dentists.PATCH(":id", dentistHandler.Patch())
			dentists.DELETE(":id", dentistHandler.Delete())
		}
		patients := api.Group("/patients")
		{
			patients.GET("", patientHandler.GetAll())
			patients.GET(":id", patientHandler.GetByID())

			patients.POST("", patientHandler.Post())
			patients.PUT(":id", patientHandler.Put())
			patients.PATCH(":id", patientHandler.Patch())
			patients.DELETE(":id", patientHandler.Delete())
		}
	}

	r.Run(fmt.Sprintf("%s", os.Getenv("HOST")))
}
