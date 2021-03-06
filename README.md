# b24imgproxy
## _Прокси трекер-пикселя писем Битрикс24_


![Build Status](https://github.com/SeverCode/b24imgproxy/actions/workflows/build.yml/badge.svg)

b24imgproxy это небольшой прокси-сервер на Go для трекер-пикселей в письмах из Битрикс24.

b24img is small proxy server for email tracking pixels sent by Bitrix24.

## Зачем?
Битрикс24 в коробочной версии зачастую находится внутри локальной сети и не имеет доступа из вне.
Однако, необходимость отослеживать статус доставки писем имеется.
Для этого случая и создано данное приложение.

This application is useful when Bitrix24 on premise version is not available from Internet.
You can use this proxy to track outgoing email statuses.   

## Особенности / Features
- Логирование запросов / Request logging
- Поддержка Docker & Docker Compose для микросервисноой архитектуры / Supports Docker & Docker Compose
- Конфигурирование параметров в YAML / YAML config file

## Запуск / Install

```sh
git clone https://github.com/SeverCode/b24imgproxy/
cd b24imgproxy
docker-compose up
```

## Самостоятельная сборка / Build and Install the Binaries from Source

```sh
git clone https://github.com/SeverCode/b24imgproxy/
cd b24imgproxy
mv config.yaml.example config.yaml && nano config.yaml
go build
```

## Как изменить URL картинки для работы через b24imgproxy
Необходимо описать функцию custom_mail в bitrix/php_interface/init.php либо local/php_interface/init.php в зависимости от того где хранятся ваши обработчики.
Код получится примерно следующий:
```php
function custom_mail($to, $subject, $message, $additional_headers = '', $additional_parameters = '')
    {
        $message = str_replace("https://bitrix24.website/pub/mail/read.php?tag=", "https://bitrix24imgproxy.website/pub/mail/read.php?tag=", $message);
        if ($additional_parameters) {
            return @mail($to, $subject, $message, $additional_headers, $additional_parameters);
        } else {
            return @mail($to, $subject, $message, $additional_headers);
        }
    }
```
* bitrix24.website - адрес вашего Битрикс24,
* bitrix24imgproxy.website - адрес прокси. 


## Подмена адреса сервера в Bitrix24

## Зависимости / Dependecies
| Dependency | License |
|-----------|----------|
|  [httprouter](https://github.com/julienschmidt/httprouter) | BSD 3-Clause "New" or "Revised" License |
| [logrus](https://github.com/sirupsen/logrus) | MIT |
| [viper](https://github.com/spf13/viper) | MIT |
| [cast](https://github.com/spf13/cast) | MIT |


## License

MIT


        