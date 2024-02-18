
# Название: **server_calculator**
# Описание: 
1) Вычисляет пример
2) Можно посмотреть ход решения примера

# Графический интерфес: **нет**

# Запуск:
1) Скачать **Zip** файл и разархивировать его
2) Открыть папку в среде разработке (у меня **Vs Code**)
3) В терминале нужно ввести команду, которая установит библиотеки
+ go mod download
4) Дальше следует запустить сервер **MySql**
Я использовал приложение [MAMP ссылка на скачивание](https://www.mamp.info/en/downloads/)
Можно посмотреть видео, как запустить [сервер My Sql](https://www.youtube.com/watch?v=4Wf__mTxm8M)

## Краткое описание
1. Нажмите start servers
2. Open web start
3. Вас перебросит на сервер
4. Дальше найдите **phpMyAdmin**
5. Вы уже зашли.
## Дальше очень важный пункт
6. Слева необходимо будет нажать на кнопку **+New** и написать **golang**
Если будете использовать другое приложение для запуска MySql,
нужно будет поменять данный в соотвествии с ваши сервером

6) Запускаем **cmd/pkg/main.go** (запускаем само приложение)
7) Заходим в консоль и пишем следующую команду
+ curl http://127.0.0.1:8080/calculate/?example=8-3*5
example указывается пример **также вместо знака + следует вводить %2B**

# Пример:
+ curl http://127.0.0.1:8080/calculate/?example=8-3%2B5
Будет возвращаться решение примера, а также id самого примера
по id можно смотреть, как решался пример введя след. команду
+ curl http://127.0.0.1:8080/steps/?id=3

# Дополнительные примеры
+ curl http://127.0.0.1:8080/calculate/?example=5-3*6/3-2
+ curl http://127.0.0.1:8080/calculate/?example=3*3%2B5-2

# Грехи проекта
Не может решать примеры со скобками
Решение с целыми числами
Деление целочислено

# Прошу прощения за ошибки писал в спешке
