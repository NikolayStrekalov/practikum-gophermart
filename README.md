# Gophermart

Учебный проект.

# Начало работы

1. Установить Go, Docker, jq, lefthook.
1. Склонируйте репозиторий в любую подходящую директорию на вашем компьютере. Синхронизируйте lefthook

      lefthook install

1. Запустить сервер:

      cd cmd/gophermart && go run main.go

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
