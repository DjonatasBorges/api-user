CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(11) NOT NULL,
    password VARCHAR(255) NOT NULL
);


INSERT INTO users (name, lastname, cpf, email, phone, password)
VALUES ('Maria', 'Souza', '25430334081', 'maria@example.com', '888888888', 'djow1234');

INSERT INTO users (name, lastname, cpf, email, phone, password)
VALUES ('Pedro', 'Ferreira', '44158277051', 'pedro@example.com', '777777777', 'djow1234');

INSERT INTO users (name, lastname, cpf, email, phone, password)
VALUES ('João', 'Silva', '31947958054', 'joao@example.com', '999999999', 'djow1234');
