# Проект "Попутчик"

## Описание проекта

**"Попутчик"** — это сервис для поиска и заказа такси, который позволяет пользователям находить попутчиков для совместных поездок. Проект направлен на оптимизацию транспортных расходов клиентов и снижение нагрузки на городскую транспортную сеть.

## Функциональные возможности
- **Роли пользователей:**
  - пассажир;
  - водитель;
  - администратор.
- **Регистрация пользователей:**
  - Регистрация пользователей реализуется с помощью роута:
    ```go
      auth.POST("/sign-up", SignUp)
    ```
- **Вход полӣзователей:**
  - Вход пользователей реализуется с помощью роута:
    ```go
      auth.POST("/sign-in", SignIn)
    ```
- **Возможность администратора:**
  - Блокировка клиентов и водителей
  - Просмотр отчёта
  - Просмотр всех маршрутов (по отклику и цене)
  - Просмотр всех маршрутов по ID (по отклику и цене)
  - Отклик маршрутов
  - Удаление маршрутов (soft delete)
  - Добавление такси компаний 
  - Просмотр всех такси компаний
  - Просмотр всех такси компаний по ID
  - Обновление такси компаний
  - Удаление такси компаний (soft delete)
    ...
- **Возможность водителя:**
  - Создание маршрутов
  - Просмотр всех маршрутов (по отклику и цене)
  - Просмотр всех маршрутов по ID (по отклику и цене)
  - Обновление маршрутов
  - Отклик или отказ маршрутов
  - Удаление маршрутов (soft delete)
  - Получение рейтинга всех пользователей
    ...
- **Возможность пользователя:**
  - Создание маршрутов
  - Просмотр всех маршрутов (по отклику и цене)
  - Просмотр всех маршрутов по ID (по отклику и цене)
  - Отклик маршрутов
  - Получение рейтинга всех водителей
    ...
## Технологии
### Backend:
- **Язык программирования:** Go (Golang)
- **База данных:** PostgreSQL
- **Фреймворки:**
  - GORM для работы с базой данных.
  - Gin для создания REST API.
- **Инструмент для проверки API:**
  - Swagger 2.0

### Дополнительно:
- **Логирование:** logrus для регистрации действий и ошибок в системе.
- **Аутентификация:** JWT-токены для безопасного доступа к API.

## Установка и запуск

### Требования:
- Go 1.19+
- PostgreSQL 16+

### Шаги по установке:
1. Клонируйте репозиторий:
    ```bash
    git clone https://github.com/n-aslik/Taxi-service.git
    ```
2. Настройте базу данных PostgreSQL:
    ```bash
    createdb test1
    ```
3. Запустите сервер:
    ```bash
    go run main.go
    ```
4. Обновить swagger:
    ```bash
    swag init
    ```


## API Документация
### API маршруты:
```go
  router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
  apiG := router.Group("/api", checkUserAuthentication)
	usersG := apiG.Group("/users")
	{
		usersG.POST("", CreateUsers)
		usersG.GET("", PrintUsers)
		usersG.GET("/:id", PrintUsersByID)
		usersG.PUT("/:id", EditUsers)
		usersG.PATCH("/:id", EditUsersRating)
		usersG.DELETE("/:id", DeleteUsers, BlockUsers)
  }
	routesG := apiG.Group("/routes")
	{
		routesG.POST("", CreateRoute)
		routesG.GET("", GetAllRoutes)
		routesG.GET("/:id", GetAllRoutesByID)
		routesG.PUT("/:id", UpdateRouteByID)
		routesG.PATCH("/:id", ChecksRouteasResponse)
		routesG.DELETE("/:id", DeleteRouteByID)
	}
	taxicompsG := apiG.Group("/taxicomps")
	{
		taxicompsG.POST("", CreateTaxicomp)
		taxicompsG.GET("", GetAllTaxiComp)
		taxicompsG.GET("/:id", GetAllTaxiCompByID)
		taxicompsG.PUT("/:id", UpdateTaxiCompByID)
		taxicompsG.DELETE("/:id", DeleteTaxiCompByID)
	}
	reportG := apiG.Group("/report")
	{
		reportG.GET("", Report)
	}
```
## Контакты
- **Разработчик:** Набиев Аслиддин
- **Email:** aslnabiev2002@gmail.com
