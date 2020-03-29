Задача:

Программа читает из stdin строки, содержащие URL. На каждый URL нужно отправить HTTP-запрос методом GET
и посчитать кол-во вхождений строки "Go" в теле ответа. В конце работы приложение выводит на экран общее количество найденных строк "Go" во всех переданных URL.

Каждый URL должен начать обрабатываться сразу после считывания и параллельно со считыванием следующего. URL должны обрабатываться параллельно, но не более k=5 одновременно. Обработчики URL не должны порождать лишних горутин, т.е. если k=5, а обрабатываемых URL-ов всего 2, не должно создаваться 5 горутин.

Пример запуска:

$ echo -e 'https://golang.org\nhttps://google.com\nhttps://yandex.ru\nhttps://mail.ru' | go run prog.go

Count for https://golang.org: 20
Count for https://google.com: 7
Count for https://yandex.ru: 4
Count for https://mail.ru: 0
Total: 31
