# Gophermart

Учебный проект.

# Начало работы

1. Установить Go, Docker, jq, lefthook.
1. Склонируйте репозиторий в любую подходящую директорию на вашем компьютере. Синхронизируйте lefthook

      lefthook install

1. Запустить базу и accrual сервис:

      docker compose up -d

1. Запустить сервер:

      cd cmd/gophermart && go run main.go

## Запуск юнит тестов

      make test

## Запуск всех тестов

      make test-all

## Создание файлов миграции

      docker run --rm \
            -v $(realpath ./internal/db/migrations):/migrations \
            migrate/migrate:v4.16.2 \
                  create \
                  -dir /migrations \
                  -ext .sql \
                  <migration_name>

## Откат миграции

      docker run --rm --network host -v $(realpath ./internal/db/migrations):/migrations \
            migrate/migrate:v4.16.2 -path=/migrations/ \
            -database  "postgres://gophermart:gophermart@localhost:5432/gophermart?sslmode=disable" \
            down -all

# Обновление шаблона

Чтобы иметь возможность получать обновления автотестов и других частей шаблона, выполните команду:

```
git remote add -m master template https://github.com/yandex-praktikum/go-musthave-diploma-tpl.git
```

Для обновления кода автотестов выполните команду:

```
git fetch template && git checkout template/master .github
```

Затем добавьте полученные изменения в свой репозиторий.
