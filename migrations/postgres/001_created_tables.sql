CREATE TABLE IF NOT EXISTS Projects (
    Id SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS Goods (
    Id SERIAL PRIMARY KEY,
    Project_id INT NOT NULL REFERENCES Projects(Id),
    Name VARCHAR(255) NOT NULL,
    Description VARCHAR(255),
    Priority INT NOT NULL,
    Removed BOOLEAN DEFAULT FALSE,
    Created_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT unique_project_goods UNIQUE (Project_id, Name)
);
