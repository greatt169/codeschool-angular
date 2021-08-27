### Гребаные кошки

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