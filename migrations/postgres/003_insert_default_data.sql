-- При добавлении записи в таблицу устанавливаем приоритет как макс приоритет в таблице +1
-- Приоритеты начинаются с 1
INSERT INTO Projects (Name) VALUES ('Первая запись') RETURNING Id, Name, Created_at;

-- SQL для вставки данных в таблицу Goods
