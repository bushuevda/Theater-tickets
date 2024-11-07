BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "Locality" (
	"id_locality"	INTEGER,
	"name"	TEXT,
	PRIMARY KEY("id_locality" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "Position" (
	"id_position"	INTEGER,
	"full_name"	TEXT,
	"short_name"	TEXT,
	PRIMARY KEY("id_position" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "Type_performance" (
	"id_type_performance"	INTEGER,
	"name"	TEXT,
	PRIMARY KEY("id_type_performance" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "Theatre" (
	"id_theatre"	INTEGER,
	"name"	TEXT,
	"id_locality"	INTEGER,
	PRIMARY KEY("id_theatre" AUTOINCREMENT),
	FOREIGN KEY("id_locality") REFERENCES "Locality"("id_locality")
);
CREATE TABLE IF NOT EXISTS "Physical_person" (
	"id_physical_person"	INTEGER,
	"surname"	TEXT,
	"name"	TEXT,
	"patronymic"	TEXT,
	"date_birthday"	TEXT,
	PRIMARY KEY("id_physical_person" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "Hall" (
	"id_hall"	INTEGER,
	"name"	TEXT,
	"id_theatre"	INTEGER,
	PRIMARY KEY("id_hall" AUTOINCREMENT),
	FOREIGN KEY("id_theatre") REFERENCES "Theatre"("id_theatre")
);
CREATE TABLE IF NOT EXISTS "Row" (
	"id_row"	INTEGER,
	"number"	INTEGER,
	"id_hall"	INTEGER,
	PRIMARY KEY("id_row" AUTOINCREMENT),
	FOREIGN KEY("id_hall") REFERENCES "Hall"("id_hall")
);
CREATE TABLE IF NOT EXISTS "Place" (
	"id_place"	INTEGER,
	"number"	INTEGER,
	"id_row"	INTEGER,
	PRIMARY KEY("id_place" AUTOINCREMENT),
	FOREIGN KEY("id_row") REFERENCES "Row"("id_row")
);
CREATE TABLE IF NOT EXISTS "Measure_unit" (
	"id_measure_unit"	INTEGER,
	"full_name"	TEXT,
	"short_name"	TEXT,
	PRIMARY KEY("id_measure_unit" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "Employment_contract" (
	"id_employment_contract"	INTEGER,
	"date_begin"	TEXT,
	"date_end"	TEXT,
	"salary"	REAL,
	"id_physical_person"	INTEGER,
	"id_position"	INTEGER,
	"id_theatre"	INTEGER,
	PRIMARY KEY("id_employment_contract" AUTOINCREMENT),
	FOREIGN KEY("id_position") REFERENCES "Position"("id_position"),
	FOREIGN KEY("id_theatre") REFERENCES "Theatre"("id_theatre")
);
CREATE TABLE IF NOT EXISTS "Performance" (
	"id_performance"	INTEGER,
	"name"	TEXT,
	"duration"	NUMERIC,
	"id_type_performance"	INTEGER,
	"id_measure_unit"	INTEGER,
	PRIMARY KEY("id_performance" AUTOINCREMENT),
	FOREIGN KEY("id_type_performance") REFERENCES "Type_performance"("id_type_performance"),
	FOREIGN KEY("id_measure_unit") REFERENCES "Measure_unit"("id_measure_unit")
);
CREATE TABLE IF NOT EXISTS "List_performance" (
	"id_list_performance"	INTEGER,
	"date"	TEXT,
	"id_theatre"	INTEGER,
	"id_performance"	INTEGER,
	PRIMARY KEY("id_list_performance" AUTOINCREMENT),
	CONSTRAINT "FK_performance" FOREIGN KEY("id_performance") REFERENCES "Performance"("id_performance"),
	CONSTRAINT "FK_theatre" FOREIGN KEY("id_theatre") REFERENCES "Theatre"("id_theatre")
);
CREATE TABLE IF NOT EXISTS "Ticket" (
	"id_ticket"	INTEGER UNIQUE,
	"price"	REAL,
	"date_performance"	TEXT,
	"time_begin"	TEXT,
	"date_sale"	TEXT,
	"id_performance"	INTEGER,
	"id_physical_person"	INTEGER,
	"id_employment_contract"	INTEGER,
	"id_place"	INTEGER,
	"id_measure_unit"	INTEGER,
	PRIMARY KEY("id_ticket" AUTOINCREMENT),
	FOREIGN KEY("id_measure_unit") REFERENCES "Measure_unit"("id_measure_unit"),
	FOREIGN KEY("id_place") REFERENCES "Place"("id_place"),
	FOREIGN KEY("id_physical_person") REFERENCES "Physical_person"("id_physical_person"),
	FOREIGN KEY("id_performance") REFERENCES "Performance"("id_performance"),
	FOREIGN KEY("id_employment_contract") REFERENCES "Employment_contract"("id_employment_contract")
);
INSERT INTO "Locality" ("id_locality","name") VALUES (1,'Оренбург'),
 (2,'Самара'),
 (3,'Москва'),
 (4,'Санкт-Петербург'),
 (5,'Пермь'),
 (6,'Томск'),
 (7,'Новосибирск'),
 (8,'Нижний Новгород'),
 (9,'Сочи'),
 (10,'Саратов');
INSERT INTO "Position" ("id_position","full_name","short_name") VALUES (1,'two','kno');
INSERT INTO "Type_performance" ("id_type_performance","name") VALUES (1,'Водевиль'),
 (2,'Драма'),
 (3,'Комедия'),
 (4,'Мелодрама'),
 (5,'Мим'),
 (6,'Моралите'),
 (7,'Мюзикл'),
 (8,'Трагедия'),
 (9,'Фарс'),
 (10,'Феерия');
INSERT INTO "Theatre" ("id_theatre","name","id_locality") VALUES (1,'Оренбургский государственный областной драматический театр им. М. Горького',1),
 (2,'Оренбургский государственный областной театр музыкальной комедии',1),
 (3,'Оренбургский государственный областной театр кукол',1),
 (4,'Оренбургский государственный татарский драматический театр им. М. Файзи',1),
 (5,'Оренбургский театр кукол «Пьеро»',1),
 (6,'Самарский академический театр оперы и балета имени Д.Д. Шостаковича',2),
 (7,'Самарский академический театр драмы им. М. Горького',2),
 (8,'Самарский театр юного зрителя «СамАрт»',2),
 (9,'Театр «Самарская площадь»',2),
 (10,'Самарский театр кукол',2),
 (11,'Театр драмы «Камерная сцена»',2),
 (12,'Самарский театр «Город»',2),
 (13,'Театр кукол «Лукоморье»',2),
 (14,'Театр «Место действия»',2),
 (15,'Театральное пространство «Дом»',2),
 (16,'Государственный академический Большой театр России (Историческая сцена)',3),
 (17,'Театр им. Евгения Вахтангова
',3),
 (18,'Московский художественный академический театр им. М. Горького',3),
 (19,'Государственный академический Малый театр России. Основная сцена',3),
 (20,'Московский государственный театр «Ленком»',3),
 (21,'Мариинский театр (Основная сцена)',4),
 (22,'Александринский театр',4),
 (23,'Большой драматический театр им. Г.А. Товстоногова',4),
 (24,'Пермский театр оперы и балета им. П.И. Чайковского',5),
 (25,'Пермский театр «У Моста»',5),
 (26,'Томский областной театр драмы',6),
 (27,'НОВАТ — Новосибирский академический театр оперы и балета. Большая сцена',7),
 (28,'Нижегородский государственный академический театр драмы им. М. Горького',8),
 (29,'Зимний театр',9),
 (30,'Саратовский академический театр оперы и балета',10);
INSERT INTO "Physical_person" ("id_physical_person","surname","name","patronymic","date_birthday") VALUES (1,'Сергеев','Иван','Сергеевич','2011-01-13');
INSERT INTO "Hall" ("id_hall","name","id_theatre") VALUES (1,'bbb',4),
 (2,'ccccbbb',24);
INSERT INTO "Row" ("id_row","number","id_hall") VALUES (1,4,7);
INSERT INTO "Place" ("id_place","number","id_row") VALUES (2,5,11),
 (3,1,1);
INSERT INTO "Measure_unit" ("id_measure_unit","full_name","short_name") VALUES (1,'рубль','руб.'),
 (2,'доллар','$'),
 (3,'ss','dsf'),
 (4,'ss','dsf'),
 (5,'aaaass','eeeeedsf');
INSERT INTO "Employment_contract" ("id_employment_contract","date_begin","date_end","salary","id_physical_person","id_position","id_theatre") VALUES (2,'2011-01-05','2011-01-07',552.55,6,6,2),
 (3,'2011-01-05','2011-01-07',552.55,6,6,3),
 (4,'2011-01-05','2011-01-07',552.55,1,2,3),
 (5,'2011-01-05','2011-01-07',552.55,6,6,5);
COMMIT;
