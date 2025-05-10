package bootstrap

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"api-fiber/app/modules/cursos"
	estudiantecurso "api-fiber/app/modules/estudiante_cursos"
	"api-fiber/database/connections"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"    // <- Importa CORS
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func InitializeApp() *fiber.App {
	// Inicialización de la BD
	dbPool, queries, err := connections.InitDB()
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	// Cleanup al cerrar la app
	cleanup := func() {
		dbPool.Close()
		log.Println("✅ Pool de conexiones a la base de datos cerrado correctamente")
	}

	// Capturar señales para shutdown limpio
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cleanup()
		os.Exit(0)
	}()

	// Cargar .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No se pudo cargar .env; usando variables de entorno del sistema")
	}

	log.Println("📌 APP_ENV:", os.Getenv("APP_ENV"))

	// Crear app Fiber
	app := fiber.New(fiber.Config{
		AppName: "Aplicación de Cursos y Estudiantes",
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(recover.New())

	// ⚡ Middleware CORS para cualquier origen, método y cabecera
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders:  "*",
		ExposeHeaders: "*",
	}))

	// Ruta de prueba
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Rutas de tus módulos
	cursos.SetupCursoRoutes(app, queries)
	estudiantecurso.SetupEstudianteCursoRoutes(app, queries)

	fmt.Println("✅ Aplicación iniciada correctamente")
	return app
}
