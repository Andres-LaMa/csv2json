# CSV ↔ JSON Converter (Go)

![Go Version](https://img.shields.io/github/go-mod/go-version/your/repo)
![CI Status](https://img.shields.io/github/actions/workflow/status/your/repo/go.yml)
![Coverage](https://img.shields.io/codecov/c/github/your/repo)
![License](https://img.shields.io/badge/license-MIT-blue)

Конвертер CSV в JSON и обратно с поддержкой:
✅ Потокового чтения/записи (для больших файлов)  
✅ Вложенных JSON-структур  
✅ Автоматического определения типов данных  
✅ Логгинга и обработки ошибок  

## Оглавление

1. [Установка](#установка)
2. [Использование](#использование)
   - [CSV → JSON](#csv--json)
   - [JSON → CSV](#json--csv)
   - [Вложенные структуры](#вложенные-структуры)
3. [Примеры](#примеры)
4. [Логгинг](#логгинг)
5. [Тестирование](#тестирование)
6. [Сборка](#сборка)
7. [Лицензия](#лицензия)

## Установка

```bash
# Клонирование репозитория
git clone https://github.com/your/repo
cd csv2json

# Сборка
make build
или
go install github.com/your/repo/cmd/converter@latest
```
##  Использование
1. CSV → JSON
    ```bash
    ./csv2json -input data.csv -output data.json
    ```
    Флаги:

    ```-input``` - входной CSV-файл

    ```-output``` - выходной JSON-файл

    ```-mode csv2json``` (по умолчанию)

2. JSON → CSV
    ```bash
    ./csv2json -mode json2csv -input data.json -output data.csv
    ```
3. Вложенные структуры

Конвертер автоматически обрабатывает:

- Вложенные объекты в JSON → преобразуются в JSON-строки в CSV

- JSON-строки в CSV → преобразуются в объекты/массивы в JSON

## Примеры

### Пример 1: Простой CSV → JSON

#### Вход (data.csv):
```
name,age
Alice,25
Bob,30
```
#### Команда:
```
./csv2json -input data.csv -output data.json
```
#### Результат (data.json):
```
[
  {
    "name": "Alice",
    "age": 25
  },
  {
    "name": "Bob",
    "age": 30
  }
]
```
### Пример 2: Вложенный JSON → CSV

#### Вход (data.json):

```
[
  {
    "name": "Alice",
    "address": {
      "city": "Moscow",
      "street": "Lenina"
    }
  }
]
```
#### Команда:
```
./csv2json -mode json2csv -input data.json -output data.csv
```

#### Результат (data.csv):
```
name,address
Alice,"{""city"":""Moscow"",""street"":""Lenina""}"
```
## Логгинг

Уровни логгинга:

- [INFO] - информация о процессе

- [ERROR] - критические ошибки

### Пример вывода:

```
[INFO] Начата конвертация data.csv → data.json
[ERROR] Ошибка чтения строки 42: неверный формат числа
[INFO] Конвертация завершена (обработано 100 строк)
```

## Тестирование
### Запуск всех тестов:
```
make test
```
### Покрытие кода:
```
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Сборка
### Доступные команды:
```
make build    # Сборка бинарника
make test     # Запуск тестов
make clean    # Очистка артефактов
```

### Сборка с версией
```
make build VERSION=1.0.0
```
## Лицензия
MIT License. Подробнее см. в файле [LICENSE](https://github.com/Andres-LaMa/csv2json/blob/main/LICENSE).

