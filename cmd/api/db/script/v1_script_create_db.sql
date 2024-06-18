

CREATE TABLE customer (
    id serial PRIMARY KEY not null,
    fullname varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    password char(64) NOT NULL,
    socialname varchar(100)
);

-- Índice para a coluna id_user
CREATE unique INDEX idx_customer_id ON customer (id);

-- Índice para a coluna email_user
CREATE unique INDEX idx_customer_email ON customer (email);

---
CREATE TABLE incomes (
    id serial PRIMARY KEY not null,
    user_id integer references customer(id),
    title varchar(50) NOT NULL,
    description varchar(100) NOT NULL,
    type varchar(50) NOT NULL,
    value decimal NOT NULL,
    datetime timestamp NOT NULL 
);

-- Índice para a coluna id
CREATE unique INDEX idx_incomes_id ON incomes (id);


-- Índice para a coluna id_user
CREATE INDEX idx_incomes_user_id ON incomes (user_id);

---

CREATE TABLE expenses (
    id serial PRIMARY KEY not null,
    user_id integer references customer(id),
    title varchar(50) NOT NULL,
    description varchar(100) NOT NULL,
    type varchar(50) NOT NULL,
    value decimal NOT NULL,
    datetime timestamp NOT NULL 
);

-- Índice para a coluna id

CREATE unique INDEX idx_expenses_id ON expenses (id);

-- Índice para a coluna id_user
CREATE INDEX idx_expenses_user_id ON expenses (user_id);