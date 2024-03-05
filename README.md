Документация

Модуль test_go

Объявление модуля:
Модуль test_go версия Go 1.16.

Зависимости:
Для этого модуля указаны конкретные версии внешних пакетов:
github.com/go-redis/redis/v8 v8.10.0
github.com/gorilla/mux v1.8.0
github.com/jackc/pgx/v4 v4.14.1
github.com/nats-io/nats-server/v2 v2.10.11 (косвенная зависимость)
github.com/nats-io/nats.go v1.33.0
Внесена замена версии github.com/nats-io/nats.go на v1.11.1-0.20210825181159-53cc936f74e9.

main.go - точка входа

Объявление пакета:
Файл main.go принадлежит пакету main.

Импорты:
В коде подключены необходимые пакеты, включая стандартные библиотеки (log, net/http) и пользовательские пакеты (test_go/handlers, test_go/utils).

Функция main:
Функция main служит точкой входа в приложение.
Она подключается к базе данных PostgreSQL, используя функцию utils.ConnectToPostgres, и регистрирует фатальную ошибку при неудаче подключения.
Создается маршрутизатор Gorilla Mux (r) для обработки HTTP-запросов.

Регистрация обработчиков:
Вызывается функция handlers.RegisterRoutes(r) для регистрации всех необходимых маршрутов и обработчиков.
Замечается комментарий, предлагающий использовать эту строку вместо handlers.RegisterGoodsHandlers(r).

Запуск HTTP-сервера:
Запускается HTTP-сервер на порту :8080 с использованием http.ListenAndServe.
Регистрируется фатальная ошибка, если сервер не запускается.

Пакет handlers
routes.go

Функциональность:
В файле реализована регистрация маршрутов с использованием Gorilla Mux в функции RegisterRoutes.
Маршрутизация:
Маршрутизатор настроен на использование обработчиков, связанных с товарами (RegisterGoodsHandlers).
Ошибка:
В комментарии указана ошибка "undefined: RegisterGoodsHandlers", что может указывать на потенциальную проблему в коде.

Пакет utils
logging.go
Функциональность:
Определена функция LogToClickhouse для отправки логов в ClickHouse через NATS.
postgres.go

Подключение к базе данных:
Функции для подключения к и закрытия соединения с базой данных PostgreSQL (ConnectToPostgres, ClosePostgres).
Функция ExecuteSQL для выполнения SQL-запросов.
redis.go
Подключение к Redis:
Функции для инициализации и закрытия соединения с сервером Redis (InitRedis, CloseRedis).
Обработка ошибок:
В коде используется log.Fatal для регистрации ошибок. Рекомендуется вернуть ошибки для более гибкой обработки.

Известная проблема:
Есть комментарий, указывающий на ошибку: handlers/routes.go:10:5: undefined: RegisterGoodsHandlers. Это говорит о том, что RegisterGoodsHandlers не определена, что может быть ошибкой или упущением.

Для устранения проблемы:
Заменить комментарий handlers.RegisterGoodsHandlers(r) на handlers.RegisterRoutes(r) в файле main.go для правильной регистрации обработчиков.

Запуск кода:
go run main.go. 


Запуск сервера Redis на Kali Linux:

Установка Redis:
sudo apt update sudo apt install redis-server 
Проверка статуса Redis:
sudo systemctl status redis-server 
Вручную запустить Redis (при необходимости):
sudo systemctl start redis-server 
Проверка работы Redis с использованием redis-cli:
redis-cli 
В интерактивной оболочке Redis выполните команды, например:
set mykey "Hello Redis" get mykey 
Остановка сервера Redis:
bash
sudo systemctl stop redis-server 
Проверка статуса Redis (должен быть inactive):
sudo systemctl status redis-server 
Отключение автозапуска Redis при загрузке системы:
sudo systemctl disable redis-server

Установка ClickHouse на Kali Linux с использованием пакетов:

Добавление репозитория ClickHouse:
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv E0C56BD4 sudo echo "deb http://repo.clickhouse.tech/deb/stable/ main/" | sudo tee /etc/apt/sources.list.d/clickhouse.list 
Обновление пакетов и установка ClickHouse:
sudo apt-get update sudo apt-get install -y clickhouse-server clickhouse-client 
Запуск ClickHouse сервера:
sudo service clickhouse-server start 
Подключение к серверу ClickHouse:
clickhouse-client

Установка NATS на Kali Linux
Загрузка бинарного файла NATS:
wget https://github.com/nats-io/nats-server/releases/download/v2.6.1/nats-server-v2.6.1-linux-amd64.zip 
Распаковка архива:
unzip nats-server-v2.6.1-linux-amd64.zip 
Переход в каталог с распакованным содержимым:
cd nats-server-v2.6.1-linux-amd64 
Запуск NATS:
./nats-server

