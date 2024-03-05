ALTER TABLE Projects ADD PRIMARY KEY (Id);
ALTER TABLE Goods ADD PRIMARY KEY (Id);
CREATE INDEX IF NOT EXISTS idx_project_name ON Projects (Name);
CREATE INDEX IF NOT EXISTS idx_goods_project_name ON Goods (Project_id, Name);
