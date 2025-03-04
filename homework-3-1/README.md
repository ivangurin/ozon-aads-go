# Хорошая хеш-функция

## Условие задачи

Лирический герой Анатолий прочитал статью про простые хеш-функции, используемые в стандартной библиотеке его любимого языка программирования, и решил в рамках своей задачи использовать подобный подход.

Для своей хеш-функции он выбрал красивое число M=127 — одно из простых чисел Мерсенна (https://ru.wikipedia.org/wiki/Число_Мерсенна#Простые_числа_Мерсенна).

Его хеш-функция даёт 64-битное беззнаковое значение.

Он определил свою хеш-функцию на следующих множествах:

- Число: целое 64-битное число со знаком.

Хеш-функция интерпретирует биты числа со знаком в дополнительном коде (https://ru.wikipedia.org/wiki/Дополнительный_код) как беззнаковое число.

Отсюда следует, что для неотрицательных чисел хеш-код совпадает с их значением.

- Символ: печатный ASCII-символ x в диапазоне 33≤x≤126.

Значение хеш-функции для — x−33.

Например, хеш-код символа ! — это 0.

- Строка: непустая последовательность символов.

Хеш-код строки, состоящей из символов c0,…cm считается по формуле:

hash(h0,…,hm)=h0⋅M^m+h1⋅M^m−1+…+hm−1⋅M+hm, где hi — хеш символа строки.

В рекурсивной форме:

hash()=0; hash(h0,…,hm)=hash(h0,…hm−1)⋅M+hm​

При расчётах используется: арифметика по модулю (https://ru.wikipedia.org/wiki/Сравнение_по_модулю).
2^64, то есть 2^64=0.

Требуется помочь Анатолию реализовать эту функцию.

## Входные данные

В первой строке вводится натуральное число N — число значений, которые требуется захешировать. N не превышает 2^31.

Последующие N строчек имеют следующий вид:

type value

где type — тип хешируемого значения, а value — само значение.

Таблица типов:
| Название  | Тип    | Запись                                   | Пример      |
|-----------|--------|------------------------------------------|-------------|
| number    | Число  | Десятичные цифры                         | 123         |
| character | Символ | Символ x в диапазоне от 33 до 126        | q           |
| string    | Строка | Последовательность символов без пробелов | HelloWorld! |

## Выходные данные

N строк, каждая из которых содержит хеш соответствующего элемента.

## Пример теста 1

### Входные данные

```
5
character f
number 7720
number -1
character s
string amogus

```

### Выходные данные

```
69
7720
18446744073709551615
82
2134387548418

```
