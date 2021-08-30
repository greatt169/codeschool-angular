#### Сайты с задачами

1. [msql](https://proglib.io/p/sql-questions)
2. [задачи + теория](https://dou.ua/lenta/articles/sql-questions/)
3. [презентация sql](https://ppt-online.org/146233)
4. [postgres полезные функции](https://habr.com/ru/post/340460/)
5. [sql оконные функции](https://habr.com/ru/post/268983/)

### Задача 1. Гребаные кошки

#### Вводные
```
CREATE TABLE public.pets (
    id integer,
    animal character varying(255)
);

COPY public.pets (id, animal) FROM stdin;
1	кошка
2	кошка
2	собака
3	пони
4	питон
4	мышь
4	хомяк
\.
```

#### Задача
Вывести всех суотрудников, у которых нет кошек
```
select distinct t1.id from pets as t1 left join pets as t2 on t1.id = t2.id and t2.animal = 'кошка' where t2.id is null
```

### Задачи из habr-статьи. По схеме

![БД](test-base.png)

#### Задача 1
Таблица Employees. Получить список всех сотрудников из 50го отдела (department_id) с зарплатой(salary), большей 4000
```
select * from employees where department_id = 50 and salary > 4000  
```

#### Задача 2
Таблица Employees. Получить список всех сотрудников у которых в имени содержатся минимум 2 буквы 'n'
```
SELECT *
  FROM employees
 WHERE first_name LIKE '%n%n%';
```

#### Задача 3
Таблица Employees. Получить список всех сотрудников у которых длина имени больше 4 букв
```
SELECT *
  FROM employees
 WHERE first_name LIKE '%_____%';
```

#### Задача 4
Таблица Employees. Получить список всех сотрудников у которых зарплата находится в промежутке от 8000 до 9000 (включительно)
```
SELECT *
  FROM employees
 WHERE salary BETWEEN 8000 and 9000;
```
#### Задача 5
Таблица Employees. Получить список всех сотрудников у которых в имени содержится символ '%'
```
SELECT *
  FROM employees
 WHERE first_name like %\%% ESCAPE '\';
```
#### Задача 6
Таблица Employees. Получить список всех ID менеджеров
```
SELECT *
  FROM employees
 WHERE first_name like %\%% ESCAPE '\';
```
#### Задача 7
Таблица Employees. Получить список всех сотрудников у которых в имени есть буква 'b' (без учета регистра)
```
SELECT *
  FROM employees
WHERE INSTR (LOWER(first_name), 'b') > 0
 ``` 

#### Задача 8

1. Как найти дубликат записи? Опишите процесс для дублирования записей с одним и несколькими полями.

```
SELECT *
  select animal from pets group by animal having count(*) > 1

 ``` 


### SQL-ex

###### Описание БД:
Схема БД состоит из четырех таблиц:
Product(maker, model, type)
PC(code, model, speed, ram, hd, cd, price)
Laptop(code, model, speed, ram, hd, price, screen)
Printer(code, model, color, type, price)
Таблица Product представляет производителя (maker), номер модели (model) и тип ('PC' - ПК, 'Laptop' - ПК-блокнот или 'Printer' - принтер). Предполагается, что номера моделей в таблице Product уникальны для всех производителей и типов продуктов. В таблице PC для каждого ПК, однозначно определяемого уникальным кодом – code, указаны модель – model (внешний ключ к таблице Product), скорость - speed (процессора в мегагерцах), объем памяти - ram (в мегабайтах), размер диска - hd (в гигабайтах), скорость считывающего устройства - cd (например, '4x') и цена - price. Таблица Laptop аналогична таблице РС за исключением того, что вместо скорости CD содержит размер экрана -screen (в дюймах). В таблице Printer для каждой модели принтера указывается, является ли он цветным - color ('y', если цветной), тип принтера - type (лазерный – 'Laser', струйный – 'Jet' или матричный – 'Matrix') и цена - price.

###### Задача 144
Найти производителей, которые производят PC как с самой низкой ценой, так и с самой высокой.
Вывод: maker

```
select t1.maker from product as t1 join pc as t2 on t1.model = t2.model WHERE t2.price = (SELECT MAX(price) FROM pc) GROUP BY t1.maker
INTERSECT
select t1.maker from product as t1 join pc as t2 on t1.model = t2.model WHERE t2.price = (SELECT MIN(price) FROM pc) GROUP BY t1.maker
```

###### Задача 18

Найдите производителей самых дешевых цветных принтеров. Вывести: maker, price

```
select t1.maker, min(t2.price) from product as t1 left join printer as t2 on t1.model = t2.model where t2.color = 'y' and t2.price=(SELECT min(price) from printer where color = 'y') group by t1.maker
```

###### Задача 19
Для каждого производителя, имеющего модели в таблице Laptop, найдите средний размер экрана выпускаемых им ПК-блокнотов.
Вывести: maker, средний размер экрана.

```
select distinct t1.maker, avg(t2.screen) from product as t1 left join laptop as t2 on t1.model = t2.model where t2.screen is not null group by t1.maker
```

###### Задача 20

Задание: 20 (Serge I: 2003-02-13)
Найдите производителей, выпускающих по меньшей мере три различных модели ПК. Вывести: Maker, число моделей ПК.

```
select maker, count(distinct model) from product where type = 'PC' group by maker having count(distinct model) > 2
```

###### Задача 21

Найдите максимальную цену ПК, выпускаемых каждым производителем, у которого есть модели в таблице PC.
Вывести: maker, максимальная цена.

```
select distinct t1.maker, max(t2.price) from product as t1 join pc as t2 on t1.model = t2.model group by t1.maker
```

###### Задача 22
Для каждого значения скорости ПК, превышающего 600 МГц, определите среднюю цену ПК с такой же скоростью. Вывести: speed, средняя цена.

```
Select speed, avg(price) from pc group by speed having speed > 600
```